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

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		ETag:             true,
		DisableKeepalive: true,
		CaseSensitive:    true,
	})

	v1 := app.Group("/v1")
	// Single route
	v1.Get("/", handler.SingleJoke)

	// Today's joke
	v1.Get("/today", handler.TodayJoke)

	// Joke by ID
	v1.Get("/:id", handler.JokeByID)

	// Add new joke
	v1.Put("/", middleware.RequireAuth(), handler.AddNewJoke)

	// Update a joke
	v1.Patch("/:id")

	// Delete a joke
	v1.Delete("/:id")

	return app
}
