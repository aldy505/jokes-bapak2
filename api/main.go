package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	v1 "jokes-bapak2-api/app/v1"
	"jokes-bapak2-api/app/v1/platform/database"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	timeoutDefault, _ := time.ParseDuration("1m")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),

		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer sentry.Flush(2 * time.Second)

	err = database.Setup()
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  timeoutDefault,
		WriteTimeout: timeoutDefault,
		ErrorHandler: errorHandler,
	})
	app.Use(cors.New())
	app.Use(limiter.New())
	app.Use(etag.New())

	app.Mount("/v1", v1.New())

	// Start server (with or without graceful shutdown).
	if os.Getenv("ENV") == "development" {
		StartServer(app)
	} else {
		StartServerWithGracefulShutdown(app)
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	log.Println(err)
	sentry.CaptureException(err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong on our end",
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
