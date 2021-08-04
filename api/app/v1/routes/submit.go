package routes

import (
	"jokes-bapak2-api/app/v1/handler/submit"

	"github.com/gofiber/fiber/v2"
)

func Submit(app *fiber.App) *fiber.App {
	// Get pending submitted joke
	app.Get("/submit", submit.GetSubmission)

	// Add a joke
	app.Post("/submit", submit.SubmitJoke)

	return app
}
