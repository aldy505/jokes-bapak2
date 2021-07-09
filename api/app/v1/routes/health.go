package routes

import (
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/handler"
	"github.com/gofiber/fiber/v2"
)

func Health(app *fiber.App) *fiber.App {
	// Health check
	app.Get("/", handler.Health)
	return app
}
