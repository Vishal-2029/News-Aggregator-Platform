package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vishal-2029/config/db"
	"github.com/Vishal-2029/models"
	"github.com/Vishal-2029/pkg/fetcher"
	natspub "github.com/Vishal-2029/pkg/nats"
	natssub "github.com/Vishal-2029/pkg/nats"
	"github.com/Vishal-2029/utility"
)


func Init() {
	log := utility.InitLogger()
	log.Println("Starting News Aggregator Service...")

	// Setup shutdown signal listener
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Init NATS Publisher
	if err := natspub.InitPublisher("nats://localhost:4222"); err != nil {
		log.Fatal("Failed to connect to NATS (pub):", err)
	}

	// Init DB
	if err := db.ConnectDB("root:root@tcp(localhost:3306)/newsdb"); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("DB connected")

	// Start gRPC server
	go func() {
		log.Println("Starting gRPC server...")
		StartGRPCServer()
	}()

	// Start NATS Subscriber
	if err := natssub.StartSubscriber("nats://localhost:4222", handleNews); err != nil {
		log.Fatal("NATS subscriber failed:", err)
	}
	log.Println("NATS subscriber running")

	// Start Cron Job
	go func() {
		log.Println("Cron job started")
		fetcher.StartCron()
	}()

	// Wait for termination
	<-ctx.Done()
	stop()

	// Log shutdown
	log.Println("Stopped fetching news.")
	log.Println("Stopping gRPC server.")
	log.Println("Shutting down News Aggregator.")
}


func handleNews(news models.NewsItem) {
	log.Println("Subscriber received:", news.Title)
	if err := db.SaveNews(news); err != nil {
		log.Println("DB Save Error:", err)
	}
}

