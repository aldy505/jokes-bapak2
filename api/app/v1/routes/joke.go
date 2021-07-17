package routes

import (
	"jokes-bapak2-api/app/v1/handler"
	"jokes-bapak2-api/app/v1/middleware"

	"github.com/gofiber/fiber/v2"
)

func Joke(app *fiber.App) *fiber.App {
	// Single route
	app.Get("/", handler.SingleJoke)

	// Today's joke
	app.Get("/today", handler.TodayJoke)

	// Joke by ID
	app.Get("/id/:id", handler.JokeByID)

	// Count total jokes
	app.Get("/total", handler.TotalJokes)

	// Add new joke
	app.Put("/", middleware.RequireAuth(), handler.AddNewJoke)

	// Update a joke
	app.Patch("/:id", middleware.RequireAuth(), handler.UpdateJoke)

	// Delete a joke
	app.Delete("/:id", middleware.RequireAuth(), handler.DeleteJoke)

	return app
}
