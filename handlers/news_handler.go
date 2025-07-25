package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Vishal-2029/config/db"
	"github.com/Vishal-2029/pkg/fetcher"
	"github.com/Vishal-2029/models"
)

// GET /news
func GetNews(c *fiber.Ctx) error {
	news, err := db.GetAllNews()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not load news"})
	}
	return c.JSON(news)
}

// POST /refresh
func RefreshNews(c *fiber.Ctx) error {
	// Fetch fresh from upstream API
	fresh := fetcher.FetchNews(nil) // `nil` is safe—your fetcher doesn’t use the NATS conn here
	if len(fresh) == 0 {
		return c.Status(500).JSON(fiber.Map{"error": "no news fetched"})
	}

	// Save each into DB (you can add duplicate checks here)
	var saved []models.NewsItem
	for _, item := range fresh {
		if err := db.SaveNews(item); err == nil {
			saved = append(saved, item)
		}
	}

	return c.JSON(fiber.Map{
		"message":      "refreshed",
		"saved_count":  len(saved),
		"new_items":    saved,
	})
}
