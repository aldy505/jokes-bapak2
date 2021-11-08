package main

import (
	"log"
	"os"
	"os/signal"

	"context"
	"jokes-bapak2-api/core/joke"
	"jokes-bapak2-api/platform/database"
	"jokes-bapak2-api/routes"

	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"github.com/Masterminds/squirrel"
	"github.com/allegro/bigcache/v3"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Setup PostgreSQL
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicln("Unable to create pool config", err)
	}
	poolConfig.MaxConnIdleTime = time.Minute * 3
	poolConfig.MaxConnLifetime = time.Minute * 5
	poolConfig.MaxConns = 15
	poolConfig.MinConns = 4

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Panicln("Unable to create connection", err)
	}
	defer db.Close()

	// Setup Redis
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	rdb := redis.NewClient(opt)
	defer rdb.Close()

	// Setup In Memory
	memory, err := bigcache.NewBigCache(bigcache.DefaultConfig(6 * time.Hour))
	if err != nil {
		log.Panicln(err)
	}
	defer memory.Close()

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

	setupCtx, setupCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*4))
	defer setupCancel()

	err = database.Populate(db, setupCtx)
	if err != nil {
		sentry.CaptureException(err)
		log.Panicln(err)
	}

	err = joke.SetAllJSONJoke(db, setupCtx, memory)
	if err != nil {
		log.Panicln(err)
	}
	err = joke.SetTotalJoke(db, setupCtx, memory)
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
	if err := a.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
