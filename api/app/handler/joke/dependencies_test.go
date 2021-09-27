package joke_test

import (
	v1 "jokes-bapak2-api/app"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App = v1.New()
