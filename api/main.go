package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	v1 "jokes-bapak2-api/app/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	timeoutDefault, _ := time.ParseDuration("1m")

	app := fiber.New(fiber.Config{
		ReadTimeout:  timeoutDefault,
		WriteTimeout: timeoutDefault,
	})

	app.Use(limiter.New(limiter.Config{
		Max:          30,
		Expiration:   1 * time.Minute,
		LimitReached: limitHandler,
	}))
	app.Use(favicon.New(favicon.Config{
		File: "./favicon.png",
	}))

	app.Mount("/v1", v1.New())

	// Start server (with or without graceful shutdown).
	if os.Getenv("ENV") == "development" {
		StartServer(app)
	} else {
		StartServerWithGracefulShutdown(app)
	}
}

func limitHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"message": "we only allow up to 15 request per minute",
	})
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
