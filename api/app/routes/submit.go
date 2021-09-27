package routes

import (
	"jokes-bapak2-api/app/handler/submit"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func (d *Dependencies) Submit() {
	deps := submit.Dependencies{
		DB:      d.DB,
		Redis:   d.Redis,
		Memory:  d.Memory,
		HTTP:    d.HTTP,
		Query:   d.Query,
		Context: d.Context,
	}

	// Get pending submitted joke
	d.App.Get(
		"/submit",
		cache.New(cache.Config{
			Expiration: 5 * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.OriginalURL()
			},
		}),
		deps.GetSubmission)

	// Add a joke
	d.App.Post("/submit", deps.SubmitJoke)
}
