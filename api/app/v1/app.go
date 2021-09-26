package v1

import (
	"context"
	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/platform/database"
	"jokes-bapak2-api/app/v1/routes"
	"log"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/allegro/bigcache/v3"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jackc/pgx/v4/pgxpool"
)

func New() *fiber.App {
	// Setup PostgreSQL
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to create pool config", err)
	}
	poolConfig.MaxConns = 15
	poolConfig.MinConns = 2

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln("Unable to create connection", err)
	}

	// Setup Redis
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	rdb := redis.NewClient(opt)

	// Setup In Memory
	memory, err := bigcache.NewBigCache(bigcache.DefaultConfig(6 * time.Hour))
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New(fiber.Config{
		DisableKeepalive: true,
		CaseSensitive:    true,
		ErrorHandler:     errorHandler,
	})

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      os.Getenv("ENV"),
		AttachStacktrace: true,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer sentry.Flush(2 * time.Second)

	err = database.Setup(db)
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

	route := routes.Dependencies{
		DB:     db,
		Redis:  rdb,
		Memory: memory,
		HTTP:   httpclient.NewClient(httpclient.WithHTTPTimeout(10 * time.Second)),
		Query:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		App:    app,
	}
	route.Health()
	route.Joke()
	route.Submit()

	return app
}

func errorHandler(c *fiber.Ctx, err error) error {
	log.Println(err)
	sentry.CaptureException(err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong on our end",
	})
}
