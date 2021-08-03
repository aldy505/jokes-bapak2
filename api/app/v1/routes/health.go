package routes

import (
	"jokes-bapak2-api/app/v1/handler/health"

	"github.com/gofiber/fiber/v2"
)

func Health(app *fiber.App) *fiber.App {
	// Health check
	app.Get("/health", health.Health)

	return app
}
