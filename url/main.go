package main

import (
    "context"
    "crypto/md5"
    "fmt"
    "log"
    "time"
   
    "encoding/base64"
    "sync"
    "net"

    ps "github.com/kasyap1234/url-shortner-microservice/proto/stats"
    pb "github.com/kasyap1234/url-shortner-microservice/proto/url"
    "github.com/redis/go-redis/v9"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type URL struct {
    ID        uint   `gorm:"primaryKey"`
    ShortURL  string `gorm:"unique"`
    LongURL   string
    CreatedAt time.Time
}

type URLShortenerServiceServer struct {
    pb.UnimplementedURLShortenerServiceServer
    statsClient     ps.StatsServiceClient
    db             *gorm.DB
    dragonflyClient *redis.Client
    mu             sync.Mutex
}

func shortenURL(url string) string {
    if url == "" {
        return ""
    }
    hash := md5.Sum([]byte(url))
    encoded := base64.URLEncoding.EncodeToString(hash[:])
    if len(encoded) < 8 {
        return encoded
    }
    return encoded[:8]
}

func (s *URLShortenerServiceServer) ShortenURL(ctx context.Context, req *pb.ShortenURLRequest) (*pb.ShortenURLResponse, error) {
    if req.LongUrl == "" {
        return nil, fmt.Errorf("long url is empty")
    }

    shortURL := shortenURL(req.LongUrl)
    s.mu.Lock()
    defer s.mu.Unlock()

    var existingURL URL
    if err := s.db.First(&existingURL, "short_url=?", shortURL).Error; err == nil {
        return &pb.ShortenURLResponse{
            ShortUrl: existingURL.ShortURL,
        }, nil
    }

    url := URL{
        ShortURL:  shortURL,
        LongURL:   req.LongUrl,
        CreatedAt: time.Now(),
    }
    if err := s.db.Create(&url).Error; err != nil {
        return nil, err
    }

    s.dragonflyClient.Set(ctx, shortURL, req.LongUrl, 10*time.Minute)
    return &pb.ShortenURLResponse{
        ShortUrl: url.ShortURL,
    }, nil
}

func (s *URLShortenerServiceServer) ResolveURL(ctx context.Context, req *pb.ResolveURLRequest) (*pb.ResolveURLResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    // Increment access count
    _, err := s.statsClient.IncrementAccessCount(ctx, &ps.IncrementAccessCountRequest{ShortUrl: req.ShortUrl})
    if err != nil {
        log.Printf("Failed to increment access count: %v", err)
    }

    // Try cache first
    cachedURL, err := s.dragonflyClient.Get(ctx, req.ShortUrl).Result()
    if err == nil {
        log.Printf("✨ Cache HIT: Retrieved %s from Dragonfly", req.ShortUrl)
        return &pb.ResolveURLResponse{
            LongUrl: cachedURL,
        }, nil
    }

    // Cache miss - fetch from database
    log.Printf("📦 Cache MISS: Fetching %s from PostgreSQL", req.ShortUrl)
    var url URL
    if err := s.db.First(&url, "short_url=?", req.ShortUrl).Error; err != nil {
        return nil, fmt.Errorf("URL not found")
    }

    // Update cache
    s.dragonflyClient.Set(ctx, req.ShortUrl, url.LongURL, 10*time.Minute)
    log.Printf("🚀 Updated cache for %s", req.ShortUrl)

    return &pb.ResolveURLResponse{
        LongUrl: url.LongURL,
    }, nil
}

// func (s *URLShortenerServiceServer) ResolveURL(ctx context.Context, req *pb.ResolveURLRequest) (*pb.ResolveURLResponse, error) {
//     s.mu.Lock()
//     defer s.mu.Unlock()

//     // Increment access count first
//     _, err := s.statsClient.IncrementAccessCount(ctx, &ps.IncrementAccessCountRequest{ShortUrl: req.ShortUrl})
//     if err != nil {
//         log.Printf("Failed to increment access count: %v", err)
//     }

//     // Try cache
//     cachedURL, err := s.dragonflyClient.Get(ctx, req.ShortUrl).Result()
//     if err == nil {
//         log.Printf("Cache hit for URL: %s", req.ShortUrl)
//         return &pb.ResolveURLResponse{
//             LongUrl: cachedURL,
//         }, nil
//     }

//     // Cache miss - fetch from database
//     var url URL
//     if err := s.db.First(&url, "short_url=?", req.ShortUrl).Error; err != nil {
//         return nil, fmt.Errorf("URL not found")
//     }

//     // Update cache
//     s.dragonflyClient.Set(ctx, req.ShortUrl, url.LongURL, 10*time.Minute)
//     log.Printf("Cache miss - URL fetched from database: %s", req.ShortUrl)

//     return &pb.ResolveURLResponse{
//         LongUrl: url.LongURL,
//     }, nil
// }



// func (s *URLShortenerServiceServer) ResolveURL(ctx context.Context, req *pb.ResolveURLRequest) (*pb.ResolveURLResponse, error) {
//     s.mu.Lock()
//     defer s.mu.Unlock()

//     cachedURL, err := s.dragonflyClient.Get(ctx, req.ShortUrl).Result()
//     if err == nil {
//         return &pb.ResolveURLResponse{
//             LongUrl: cachedURL,
//         }, nil
//     }

//     var url URL
//     if err := s.db.First(&url, "short_url=?", req.ShortUrl).Error; err != nil {
//         return nil, fmt.Errorf("URL not found")
//     }

//     s.dragonflyClient.Set(ctx, req.ShortUrl, url.LongURL, 10*time.Minute)
//     s.statsClient.IncrementAccessCount(ctx, &ps.IncrementAccessCountRequest{ShortUrl: req.ShortUrl})

//     return &pb.ResolveURLResponse{
//         LongUrl: url.LongURL,
//     }, nil
// }

func main() {
    const urlServicePort = ":50051"

    db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    db.AutoMigrate(&URL{})

    dragonflyClient := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    statsConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("failed to connect to stats service: %v", err)
    }
    statsClient := ps.NewStatsServiceClient(statsConn)

    server := &URLShortenerServiceServer{
        statsClient:     statsClient,
        db:             db,
        dragonflyClient: dragonflyClient,
    }

    grpcServer := grpc.NewServer()
    pb.RegisterURLShortenerServiceServer(grpcServer, server)

    lis, err := net.Listen("tcp", urlServicePort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Println("URL Shortener service server started on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
