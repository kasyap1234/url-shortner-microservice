package main

import (
    "context"
    "fmt"
    "log"
    "time"

    statsPb "github.com/kasyap1234/url-shortner-microservice/proto/stats"
    urlPb "github.com/kasyap1234/url-shortner-microservice/proto/url"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

const (
    urlServiceAddr   = "localhost:50051"
    statsServiceAddr = "localhost:50052"
)

type ClientApp struct {
    urlClient   urlPb.URLShortenerServiceClient
    statsClient statsPb.StatsServiceClient
}

func NewClientApp() (*ClientApp, error) {
    urlConn, err := grpc.Dial(urlServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to URL service: %v", err)
    }
    
    statsConn, err := grpc.Dial(statsServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Stats service: %v", err)
    }
    
    return &ClientApp{
        urlClient:   urlPb.NewURLShortenerServiceClient(urlConn),
        statsClient: statsPb.NewStatsServiceClient(statsConn),
    }, nil
}

func (c *ClientApp) ShortenURL(originalUrl string) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    resp, err := c.urlClient.ShortenURL(ctx, &urlPb.ShortenURLRequest{LongUrl: originalUrl})
    if err != nil {
        return "", fmt.Errorf("failed to shorten URL: %v", err)
    }
    return resp.ShortUrl, nil
}
func (c *ClientApp) GetAccessCount(shortUrl string) (int64, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    resp, err := c.statsClient.GetAccessCount(ctx, &statsPb.GetAccessCountRequest{ShortUrl: shortUrl})
    if err != nil {
        return 0, fmt.Errorf("failed to get access count: %v", err)
    }
    return resp.AccessCount, nil
}
func (c *ClientApp) ResolveURL(shortUrl string)(string,error){
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	resp,err := c.urlClient.ResolveURL(ctx,&urlPb.ResolveURLRequest{ShortUrl:shortUrl})
	if err !=nil {
		return "",fmt.Errorf("failed to resolve URL: %v",err)
	}
	return resp.LongUrl,nil 
}

func main() {
    client, err := NewClientApp()
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    
    originalUrl := "https://longurl.com/very/long/url"
    shortUrl, err := client.ShortenURL(originalUrl)
    if err != nil {
        log.Fatalf("Failed to shorten URL: %v", err)
    }
    fmt.Printf("Successfully shortened URL: %s\n", shortUrl)

    // Make multiple resolve requests to test caching
    for i := 0; i < 5; i++ {
        start := time.Now()
        longUrl, err := client.ResolveURL(shortUrl)
        duration := time.Since(start)
        
        if err != nil {
            log.Printf("Failed to resolve URL: %v", err)
            continue 
        }
        fmt.Printf("Request %d - Resolved to: %s (took: %v)\n", i+1, longUrl, duration)
        
        count, _ := client.GetAccessCount(shortUrl)
        fmt.Printf("Current access count: %d\n", count)
        
        time.Sleep(1 * time.Second) // Small delay between requests
    }
}
