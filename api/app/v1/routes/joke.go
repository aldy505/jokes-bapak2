package routes

import (
	"jokes-bapak2-api/app/v1/handler/joke"
	"jokes-bapak2-api/app/v1/middleware"

	"github.com/gofiber/fiber/v2"
)

func Joke(app *fiber.App) *fiber.App {
	// Single route
	app.Get("/", joke.SingleJoke)

	// Today's joke
	app.Get("/today", joke.TodayJoke)

	// Joke by ID
	app.Get("/id/:id", middleware.OnlyIntegerAsID(), joke.JokeByID)

	// Count total jokes
	app.Get("/total", joke.TotalJokes)

	// Add new joke
	app.Put("/", middleware.RequireAuth(), joke.AddNewJoke)

	// Update a joke
	app.Patch("/id/:id", middleware.RequireAuth(), middleware.OnlyIntegerAsID(), joke.UpdateJoke)

	// Delete a joke
	app.Delete("/id/:id", middleware.RequireAuth(), middleware.OnlyIntegerAsID(), joke.DeleteJoke)

	return app
}
