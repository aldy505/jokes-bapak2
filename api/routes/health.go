package routes

import (
	"jokes-bapak2-api/handler/health"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

// Health provides route for healthcheck endpoints.
func Health(bucket *minio.Client, cache *redis.Client) *chi.Mux {
	dependency := &health.Dependencies{
		Bucket: bucket,
		Cache:  cache,
	}

	router := chi.NewRouter()

	router.Get("/", dependency.Health)

	return router
}
