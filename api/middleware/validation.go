package middleware

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func OnlyIntegerAsID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		regex, err := regexp.Compile(`([0-9]+)`)
		if err != nil {
			return err
		}

		loc := regex.FindStringIndex(c.Params("id"))
		if loc[1] == len(c.Params("id")) {
			return c.Next()
		}

		return c.
			Status(fiber.StatusBadRequest).
			JSON(Error{
				Error: "only numbers are allowed as ID",
			})
	}
}
