package routes

import (
	"jokes-bapak2-api/app/v1/handler/submit"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Submit(app *fiber.App) *fiber.App {
	// Get pending submitted joke
	app.Get(
		"/submit",
		cache.New(cache.Config{
			Expiration: 5 * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.OriginalURL()
			},
		}),
		submit.GetSubmission)

	// Add a joke
	app.Post("/submit", submit.SubmitJoke)

	return app
}
