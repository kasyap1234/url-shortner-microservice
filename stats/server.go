package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/kasyap1234/url-shortner-microservice/proto/stats"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type StatsServiceServer struct {
// 	pb.UnimplementedStatsServiceServer
// 	mu          sync.Mutex
// 	accessCount map[string]int64
// }
var mu sync.Mutex 
var db *gorm.DB 

// func NewStatsServiceServer() *StatsServiceServer {
// 	return &StatsServiceServer{
// 		accessCount: make(map[string]int64),
// 	}

// }

// func (s *StatsServiceServer) IncrementAccessCount(ctx context.Context, req *pb.IncrementAccessCountRequest) (*pb.IncrementAccessCountResponse, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	if req.ShortUrl == "" {
// 		return nil, fmt.Errorf("short url is required")
// 	}

// 	s.accessCount[req.ShortUrl]++
// 	log.Println("increment access for url ", req.ShortUrl)
// 	return &pb.IncrementAccessCountResponse{
// 		Success: true,
// 	}, nil

// }

// func (s *StatsServiceServer) GetAccessCount(ctx context.Context, req *pb.GetAccessCountRequest) (*pb.GetAccessCountResponse, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	if req.ShortUrl == "" {
// 		return nil, fmt.Errorf("short url is required")
// 	}
// 	count, exists := s.accessCount[req.ShortUrl]
// 	if !exists {
// 		return nil, fmt.Errorf("short url not found")
// 	}
// 	return &pb.GetAccessCountResponse{
// 		AccessCount: count,
// 	}, nil
// }

// func main(){
// 	lis,err := net.Listen("tcp",":50052")
// 	if err != nil{
// 		log.Fatalf("failed to listen: %v",err)
// 	}
// 	grpcServer := grpc.NewServer()
// 	statsServer := NewStatsServiceServer()
// 	pb.RegisterStatsServiceServer(grpcServer,statsServer)
// 	log.Println("Stats service server started on port 50052")
// 	if err := grpcServer.Serve(lis); err != nil{
// 		log.Fatalf("failed to serve: %v",err)
// 	}
// }


type Stats struct {
	ID uint `gorm:"primaryKey"`
	ShortUrl string `gorm:"unique"`
	Count  int64
}

type StatsServiceServer struct {
 pb.UnimplementedStatsServiceServer
}
func NewStatsServiceServer() *StatsServiceServer{
	return &StatsServiceServer{
	
	}
}


func (s*StatsServiceServer)IncrementAccessCount(ctx context.Context,req *pb.IncrementAccessCountRequest)(*pb.IncrementAccessCountResponse,error){
	mu.Lock()
	defer mu.Unlock()
	var stats Stats 
	if err :=db.First(&stats,"short_url = ?",req.ShortUrl).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			stats=Stats{ShortUrl: req.ShortUrl,Count: 1}

			if err:=db.Create(&stats).Error; err !=nil {
				return nil,err ; 
			}

		}
	}else{
		stats.Count++
		if err :=db.Save(&stats).Error; err != nil{
			return nil,err
		}
		return &pb.IncrementAccessCountResponse{Success: true},nil ; 

	}

func (s *StatsServiceServer) GetAccessCount(ctx context.Context, req *pb.GetAccessCountRequest)(*pb.GetAccessCountResponse,error){
	mu.Lock()
	defer mu.Unlock()
	var stats Stats
	if err :=db.First(&stats,"short_url = ?",req.ShortUrl).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil,fmt.Errorf("short url not found")
		}
		return nil,err
	}
	return &pb.GetAccessCountResponse{AccessCount: stats.Count},nil
}

func main(){
	var err error 
	db,err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"),&gorm.Config{})
	if err !=nil {

	}
	db.AutoMigrate(&Stats{})
	listener,err :=net.Listen("tcp",":50052")
	if err !=nil {
		log.Fatalf("failed to listen: %v",err)

	}
	grpcServer :=grpc.NewServer()
    statsServer := NewStatsServiceServer()
	// registering the microservice ; 
	pb.RegisterStatsServiceServer(grpcServer,statsServer)
	log.Println("Stats service server started on port 50052")
	if err := grpcServer.Serve(listener); err != nil{
		log.Fatalf("failed to serve: %v",err)
	}

}