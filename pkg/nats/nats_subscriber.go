package nats

import (
	"encoding/json"
	"log"

	"github.com/Vishal-2029/models"
	"github.com/nats-io/nats.go"
)

// StartSubscriber connects and listens to published news
func StartSubscriber(url string, handler func(models.NewsItem)) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}
	log.Println("✅ Connected to NATS (Subscriber)")

	_, err = nc.Subscribe(Subject, func(msg *nats.Msg) {
		var news models.NewsItem
		if err := json.Unmarshal(msg.Data, &news); err != nil {
			log.Println("Error decoding message:", err)
			return
		}
		handler(news)
	})
	return err
}
