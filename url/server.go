package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"

	"time"

	"strings"

	"encoding/hex"
	"sync"

	ps "github.com/kasyap1234/url-shortner-microservice/proto/stats"
	pb "github.com/kasyap1234/url-shortner-microservice/proto/url"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"net"
)

var db *gorm.DB
var dragonflyClient *redis.Client
var ctx = context.Background()
var mu sync.Mutex

// Base62 characters
const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//URL shortening ;

// Hash the URL using MD5 and convert to Base62
func shortenURL(longURL string) string {
	hash := md5.Sum([]byte(longURL))
	hexHash := hex.EncodeToString(hash[:])
	return base62Encode(hexHash)[:8] // Take the first 8 characters
}

// Base62 encoding function
func base62Encode(input string) string {
	var result strings.Builder
	num := int64(0)

	// Convert hex string to int
	for _, c := range input {
		num = num*16 + int64(strings.Index("0123456789abcdef", string(c)))
	}

	// Base62 encoding
	for num > 0 {
		result.WriteByte(base62Chars[num%62])
		num /= 62
	}

	return result.String()
}

// defining URL for postgresql

type URL struct {
	ID        uint   `gorm:"primaryKey"`
	ShortURL  string `gorm:"unique"`
	LongURL   string
	CreatedAt time.Time
}

// pb.UnimplementedURLShortenerServiceServer for covering all the bases of grpc (failures,error success responses);
type URLShortenerServiceServer struct {
	pb.UnimplementedURLShortenerServiceServer
	statsClient     ps.StatsServiceClient
	db              *gorm.DB
	dragonflyClient *redis.Client
	mu              sync.Mutex
}

// actual shortening url function

func (s *URLShortenerServiceServer) ShortenURL(ctx context.Context, req *pb.ShortenURLRequest) (*pb.ShortenURLResponse, error) {
	shortURL := shortenURL(req.LongUrl)
	mu.Lock()
	defer mu.Unlock()
	var existingURL URL
	if err := db.First(&existingURL, "short_url=?", shortURL).Error; err == nil {
		return &pb.ShortenURLResponse{
			ShortUrl: existingURL.ShortURL,
		}, nil

	}
	url := URL{
		ShortURL:  shortURL,
		LongURL:   req.LongUrl,
		CreatedAt: time.Now(),
	}
	if err := db.Create(&url).Error; err != nil {
		return nil, err
	}
	// caching the url in dragonfly db ;
	dragonflyClient.Set(ctx, shortURL, req.LongUrl, 10*time.Minute)

	return &pb.ShortenURLResponse{
		ShortUrl: url.ShortURL,
	}, nil

}

// get the long url from the short url (or resolve the url)
func (s *URLShortenerServiceServer) ResolveURL(ctx context.Context, req *pb.ResolveURLRequest) (*pb.ResolveURLResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	cachedURL, err := dragonflyClient.Get(ctx, req.ShortUrl).Result()
	if err == nil {
		return &pb.ResolveURLResponse{
			LongUrl: cachedURL,
		}, nil

	}
	var url URL
	if err := db.First(&url, "short_url=?", req.ShortUrl).Error; err != nil {
		return nil, fmt.Errorf("URL not found")

	}
	dragonflyClient.Set(ctx, req.ShortUrl, url.LongURL, 10*time.Minute)
	s.statsClient.IncrementAccessCount(ctx, &ps.IncrementAccessCountRequest{ShortUrl: req.ShortUrl})

	return &pb.ResolveURLResponse{
		LongUrl: url.LongURL,
	}, nil
}

// func incrementAccessCount(shortURL string) {
// 	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		log.Fatalf("failed to connect to stats service: %v", err)
// 		return
// 	}
// 	defer conn.Close()

// }
func main() {
	var err error
	const urlServicePort = ":50051"

	db, err = gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&URL{})
	dragonflyClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to stats service: %v", err)
	}
	statsClient := ps.NewStatsServiceClient(conn)

	server := &URLShortenerServiceServer{
		statsClient:     statsClient,
		db:              db,
		dragonflyClient: dragonflyClient,
	}
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", urlServicePort)
	if err !=nil {
		log.Fatalf("failed to listen: %v", err)

	}
	pb.RegisterURLShortenerServiceServer(grpcServer, server)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
