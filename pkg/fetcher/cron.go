package fetcher

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Vishal-2029/models"
	natspub "github.com/Vishal-2029/pkg/nats"
	"github.com/robfig/cron/v3"
)

var lastSeenIDs = make(map[string]bool)

func StartCron() {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		StartNewsJob()
	})
	c.Start()
	StartNewsJob()
}

func StartNewsJob() {
	newsList := FetchNews(nil)

	for _, item := range newsList {
		if !lastSeenIDs[item.ArticleId] {
			err := natspub.PublishNewsItem(item)
			if err != nil {
				log.Println("Error publishing to NATS:", err)
			}
			printNews(item)
			lastSeenIDs[item.ArticleId] = true
			time.Sleep(5 * time.Second) // space out publishing
		}
	}
}

func printNews(item models.NewsItem) {
	fmt.Printf("📰 Title       : %s\n", item.Title)
	fmt.Printf("🆔 Article ID  : %s\n", item.ArticleId)
	fmt.Printf("🖋️  Author      : %s\n", strings.Join(item.Creator, ", "))
	fmt.Printf("🌐 Source      : %s\n", item.Source)
	fmt.Printf("🖼️  Image URL   : %s\n", item.ImageURl)
	fmt.Printf("🔗 Link        : %s\n", item.Link)
	fmt.Printf("📝 Summary     : %s\n", item.Description)
	fmt.Printf("📅 Published   : %s\n", item.PublishedAt)
	fmt.Println("")
}

