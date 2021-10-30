package routes

import (
	"jokes-bapak2-api/handler/joke"
	"jokes-bapak2-api/middleware"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cache"
)

func (d *Dependencies) Joke() {
	deps := joke.Dependencies{
		DB:     d.DB,
		Redis:  d.Redis,
		Memory: d.Memory,
		HTTP:   d.HTTP,
		Query:  d.Query,
	}
	// Single route
	d.App.Get("/", deps.SingleJoke)
	d.App.Get("/v1", deps.SingleJoke)

	// Today's joke
	d.App.Get("/today", cache.New(cache.Config{Expiration: 6 * time.Hour}), deps.TodayJoke)
	d.App.Get("/v1/today", cache.New(cache.Config{Expiration: 6 * time.Hour}), deps.TodayJoke)

	// Joke by ID
	d.App.Get("/id/:id", middleware.OnlyIntegerAsID(), deps.JokeByID)
	d.App.Get("/v1/id/:id", middleware.OnlyIntegerAsID(), deps.JokeByID)

	// Count total jokes
	d.App.Get("/total", cache.New(cache.Config{Expiration: 15 * time.Minute}), deps.TotalJokes)
	d.App.Get("/v1/total", cache.New(cache.Config{Expiration: 15 * time.Minute}), deps.TotalJokes)

	// Add new joke
	d.App.Put("/", middleware.RequireAuth(d.DB), deps.AddNewJoke)

	// Update a joke
	d.App.Patch("/id/:id", middleware.RequireAuth(d.DB), middleware.OnlyIntegerAsID(), deps.UpdateJoke)

	// Delete a joke
	d.App.Delete("/id/:id", middleware.RequireAuth(d.DB), middleware.OnlyIntegerAsID(), deps.DeleteJoke)
}
