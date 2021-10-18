package routes

import (
	"jokes-bapak2-api/app/handler/health"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cache"
)

func (d *Dependencies) Health() {
	// Health check
	deps := health.Dependencies{
		DB:    d.DB,
		Redis: d.Redis,
	}

	d.App.Get("/health", cache.New(cache.Config{Expiration: 30 * time.Minute}), deps.Health)
	d.App.Get("/v1/health", cache.New(cache.Config{Expiration: 30 * time.Minute}), deps.Health)
}
