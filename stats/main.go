package main

import (
	"context"

	"log"
	"net"
	"sync"

	pb "github.com/kasyap1234/url-shortner-microservice/proto/stats"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Stats struct {
	ID       uint   `gorm:"primaryKey"`
	ShortUrl string `gorm:"unique"`
	Count    int64
}

type StatsServiceServer struct {
	pb.UnimplementedStatsServiceServer
	db *gorm.DB
	mu sync.Mutex
}



func (s *StatsServiceServer) IncrementAccessCount(ctx context.Context, req *pb.IncrementAccessCountRequest) (*pb.IncrementAccessCountResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var stats Stats
	if err := s.db.First(&stats, "short_url = ?", req.ShortUrl).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			stats = Stats{ShortUrl: req.ShortUrl, Count: 1}
			if err := s.db.Create(&stats).Error; err != nil {
				return nil, err
			}
		}
	} else {
		stats.Count++
		if err := s.db.Save(&stats).Error; err != nil {
			return nil, err
		}
	}
	return &pb.IncrementAccessCountResponse{Success: true}, nil
}

func main() {
	var err error

	if err != nil {

	}
	db.AutoMigrate(&Stats{})
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	grpcServer := grpc.NewServer()
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	statsServer := &StatsServiceServer{
		db: db,
	}

	// registering the microservice ;
	pb.RegisterStatsServiceServer(grpcServer, statsServer)
	log.Println("Stats service server started on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
