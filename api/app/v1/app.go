package v1

import (
	"jokes-bapak2-api/app/v1/routes"

	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableKeepalive: true,
		CaseSensitive:    true,
	})

	routes.Health(app)
	routes.Joke(app)

	return app
}
