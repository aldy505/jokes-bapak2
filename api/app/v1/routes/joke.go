package routes

import (
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/handler"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/middleware"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/platform/cache"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/platform/database"

	"github.com/gofiber/fiber/v2"
)

var db = database.New()
var redis = cache.New()

func Joke(app *fiber.App) *fiber.App {
	// Single route
	app.Get("/", handler.SingleJoke)

	// Today's joke
	app.Get("/today", handler.TodayJoke)

	// Joke by ID
	app.Get("/:id", handler.JokeByID)

	// Add new joke
	app.Put("/", middleware.RequireAuth(), handler.AddNewJoke)

	// Update a joke
	app.Patch("/:id", middleware.RequireAuth(), handler.UpdateJoke)

	// Delete a joke
	app.Delete("/:id", middleware.RequireAuth(), handler.DeleteJoke)

	return app
}
