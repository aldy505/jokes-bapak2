package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"jokes-bapak2-api/core/joke"
	"jokes-bapak2-api/routes"

	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/allegro/bigcache/v3"
	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	redisURL, ok := os.LookupEnv("REDIS_URL")
	if !ok {
		redisURL = "redis://@localhost:6379"
	}

	minioHost, ok := os.LookupEnv("MINIO_HOST")
	if !ok {
		minioHost = "localhost:9000"
	}

	minioRegion, ok := os.LookupEnv("MINIO_REGION")
	if !ok {
		minioRegion = ""
	}

	minioSecure, ok := os.LookupEnv("MINIO_SECURE")
	if !ok {
		minioSecure = "false"
	}

	minioID, ok := os.LookupEnv("MINIO_ACCESS_ID")
	if !ok {
		minioID = "minio"
	}

	minioSecret, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		minioSecret = "password"
	}

	minioToken, ok := os.LookupEnv("MINIO_TOKEN")
	if !ok {
		minioToken = ""
	}

	sentryDsn, ok := os.LookupEnv("SENTRY_DSN")
	if !ok {
		sentryDsn = ""
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}

	hostname, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		hostname = "127.0.0.1"
	}

	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		environment = "development"
	}

	// Setup In Memory
	memory, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Panicln(err)
	}
	defer memory.Close()

	// Setup MinIO
	minioClient, err := minio.New(minioHost, &minio.Options{
		Creds:  credentials.NewStaticV4(minioID, minioSecret, minioToken),
		Region: minioRegion,
		Secure: minioSecure == "true",
	})
	if err != nil {
		log.Fatalf("setting up minio client: %s", err.Error())
		return
	}

	parsedRedisURL, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("parsing redis url: %s", err.Error())
		return
	}

	redisClient := redis.NewClient(parsedRedisURL)
	defer func() {
		err := redisClient.Close()
		if err != nil {
			log.Printf("closing redis client: %s", err.Error())
		}
	}()

	// Setup Sentry
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		Environment:      environment,
		AttachStacktrace: true,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: environment != "production",
	})
	if err != nil {
		log.Fatalf("setting up sentry: %s", err.Error())
		return
	}
	defer sentry.Flush(2 * time.Second)

	setupCtx, setupCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*4))
	defer setupCancel()

	_, _, err = joke.GetTodaysJoke(setupCtx, minioClient, redisClient, memory)
	if err != nil {
		log.Fatalf("getting initial joke data: %s", err.Error())
		return
	}

	healthRouter := routes.Health(minioClient, redisClient)
	jokeRouter := routes.Joke(minioClient, redisClient, memory)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet},
		AllowCredentials: false,
		MaxAge:           int(60 * 60 * 24 * 365),
		Debug:            false,
	}).Handler)

	router.Mount("/health", healthRouter)
	router.Mount("/", jokeRouter)

	server := &http.Server{
		Handler:           router,
		Addr:              net.JoinHostPort(hostname, port),
		ReadTimeout:       time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Minute,
	}

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listening http server: %v", err)
		}
	}()

	<-exitSignal

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer shutdownCancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("shutting down http server: %v", err)
	}
}
