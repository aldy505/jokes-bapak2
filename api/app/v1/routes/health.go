package routes

import (
	"jokes-bapak2-api/app/v1/handler/health"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Health(app *fiber.App) *fiber.App {
	// Health check
	app.Get("/health", cache.New(cache.Config{Expiration: 30 * time.Minute}), health.Health)

	return app
}
