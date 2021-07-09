package main

import (
	"log"
	"os"
	"time"

	"github.com/aldy505/jokes-bapak2-api/api/app/v1/platform/database"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/routes"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	app.Mount("/", routes.New())

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func errorHandler(c *fiber.Ctx, err error) error {
	sentry.CaptureException(err)
	return c.Status(500).JSON(fiber.Map{
		"error": "Something went wrong on our end",
	})
}
