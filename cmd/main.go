package cmd

import (
	"log"


	"github.com/Vishal-2029/config/db"
	"github.com/Vishal-2029/models"
	"github.com/Vishal-2029/pkg/fetcher"
	natspub "github.com/Vishal-2029/pkg/nats"
	natssub "github.com/Vishal-2029/pkg/nats"
)

func Init() {
	// Init NATS Publisher
	err := natspub.InitPublisher("nats://localhost:4222")
	if err != nil {
		log.Fatal("Failed to connect to NATS (pub):", err)
	}

	// Init DB
	err = db.ConnectDB("root:root@tcp(localhost:3306)/newsdb")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("✅ DB connected")

	// Start gRPC server
	go StartGRPCServer()

	// Start NATS Subscriber
	err = natssub.StartSubscriber("nats://localhost:4222", handleNews)
	if err != nil {
		log.Fatal("NATS subscriber failed:", err)
	}

	// Start Cron job to fetch news
	go fetcher.StartCron()

	select {} // keep the app running
}

func handleNews(news models.NewsItem) {
	log.Println("🟢 Subscriber received:", news.Title)
	if err := db.SaveNews(news); err != nil {
		log.Println("DB Save Error:", err)
	}
}

