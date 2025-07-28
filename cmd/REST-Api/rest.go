package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Vishal-2029/config/db"
	"github.com/Vishal-2029/handlers"
	"github.com/Vishal-2029/utility"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log := utility.InitLogger()

	// Connect DB
	if err := db.ConnectDB("root:root@tcp(localhost:3306)/newsdb"); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("DB connected")

	app := fiber.New()

	// Routes
	app.Get("/news", handlers.GetNews)
	app.Post("/refresh", handlers.RefreshNews)

	// Create context to listen for OS interrupt (Ctrl+C)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Run server in background
	go func() {
		log.Println("REST API running at http://localhost:3000")
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("REST server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	stop()

	// Shutdown process
	log.Println("Shutting down REST API...")
	_ = app.Shutdown() // optional: gracefully stop Fiber

	// Wait to ensure all logs print cleanly
	time.Sleep(1 * time.Second)
	log.Println("REST API stopped.")
}
