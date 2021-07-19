package v1

import (
	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/platform/cache"
	"jokes-bapak2-api/app/v1/platform/database"
	"jokes-bapak2-api/app/v1/routes"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

var memory = cache.InMemory()
var db = database.New()

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableKeepalive: true,
		CaseSensitive:    true,
		ErrorHandler:     errorHandler,
	})

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("ENV"),
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

	err = core.SetAllJSONJoke(db, memory)
	if err != nil {
		log.Fatalln(err)
	}
	err = core.SetTotalJoke(db, memory)
	if err != nil {
		log.Fatalln(err)
	}

	app.Use(cors.New())
	app.Use(etag.New())

	routes.Health(app)
	routes.Joke(app)

	return app
}

func errorHandler(c *fiber.Ctx, err error) error {
	log.Println(err)
	sentry.CaptureException(err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong on our end",
	})
}
