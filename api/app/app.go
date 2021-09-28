package app

import (
	"context"
	"jokes-bapak2-api/app/core"
	"jokes-bapak2-api/app/platform/database"
	"jokes-bapak2-api/app/routes"
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
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func New() *fiber.App {
	// Setup Context
	ctx, cancel := context.WithCancel(context.Background())

	// Setup PostgreSQL
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicln("Unable to create pool config", err)
	}
	poolConfig.MaxConnIdleTime = time.Minute * 3
	poolConfig.MaxConnLifetime = time.Minute * 5
	poolConfig.MaxConns = 15
	poolConfig.MinConns = 4

	db, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Panicln("Unable to create connection", err)
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
		log.Panicln(err)
	}

	// Setup Sentry
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      os.Getenv("ENV"),
		AttachStacktrace: true,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		log.Panicln(err)
	}
	defer sentry.Flush(2 * time.Second)

	err = database.Setup(db, &ctx)
	if err != nil {
		sentry.CaptureException(err)
		log.Panicln(err)
	}

	err = core.SetAllJSONJoke(db, memory, &ctx)
	if err != nil {
		log.Panicln(err)
	}
	err = core.SetTotalJoke(db, memory, &ctx)
	if err != nil {
		log.Panicln(err)
	}

	timeoutDefault := time.Minute * 1

	app := fiber.New(fiber.Config{
		ReadTimeout:      timeoutDefault,
		WriteTimeout:     timeoutDefault,
		CaseSensitive:    true,
		DisableKeepalive: true,
		ErrorHandler:     errorHandler,
	})

	app.Use(limiter.New(limiter.Config{
		Max:          30,
		Expiration:   1 * time.Minute,
		LimitReached: limitHandler,
	}))

	app.Use(cors.New())
	app.Use(etag.New())

	route := routes.Dependencies{
		DB:      db,
		Redis:   rdb,
		Memory:  memory,
		HTTP:    httpclient.NewClient(httpclient.WithHTTPTimeout(10 * time.Second)),
		Query:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		App:     app,
		Context: &ctx,
		Cancel:  &cancel,
	}
	route.Health()
	route.Joke()
	route.Submit()

	return app
}

func limitHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
		"message": "we only allow up to 15 request per minute",
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	log.Println(err)
	sentry.CaptureException(err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong on our end",
	})
}
