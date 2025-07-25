package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/Vishal-2029/config/db"
	"github.com/Vishal-2029/handlers"
)

func main() {
	// Initialize only DB
	err := db.ConnectDB("root:root@tcp(localhost:3306)/newsdb")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("✅ DB connected")

	app := fiber.New()

	app.Get("/news", handlers.GetNews)
	app.Post("/refresh", handlers.RefreshNews)

	log.Println("🚀 REST API running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}