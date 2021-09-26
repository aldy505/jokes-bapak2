package routes

import (
	"jokes-bapak2-api/app/v1/handler/health"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func (d *Dependencies) Health() *fiber.App {
	// Health check
	deps := health.Dependencies{
		Redis: d.Redis,
	}
	d.App.Get("/health", cache.New(cache.Config{Expiration: 30 * time.Minute}), deps.Health)

	return d.App
}
