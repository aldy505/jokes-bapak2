package routes

import (
	"jokes-bapak2-api/handler/joke"

	"github.com/allegro/bigcache/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

// Joke provides route for jokes.
func Joke(bucket *minio.Client, cache *redis.Client, memory *bigcache.BigCache) *chi.Mux {
	deps := &joke.Dependencies{
		Memory: memory,
		Bucket: bucket,
		Redis:  cache,
	}

	router := chi.NewRouter()

	// Single route
	router.Get("/", deps.SingleJoke)

	// Today's joke
	router.Get("/today", deps.TodayJoke)

	// Joke by ID
	router.Get("/id/{id}", deps.JokeByID)

	// Count total jokes
	router.Get("/total", deps.TotalJokes)

	return router
}
