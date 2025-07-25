package nats

import (
	"encoding/json"
	"log"

	"github.com/Vishal-2029/models"
	"github.com/nats-io/nats.go"
)

const Subject = "news.items"

var natsConn *nats.Conn

// InitPublisher establishes the NATS connection
func InitPublisher(url string) error {
	var err error
	natsConn, err = nats.Connect(url)
	if err != nil {
		return err
	}
	log.Println("✅ Connected to NATS (Publisher)")
	return nil
}

// PublishNewsItem sends a single news item to the subject
func PublishNewsItem(news models.NewsItem) error {
	data, err := json.Marshal(news)
	if err != nil {
		return err
	}

	err = natsConn.Publish(Subject, data)
	if err != nil {
		return err
	}

	log.Println("🟢 Published:", news.Title)
	return nil
}
