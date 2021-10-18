package joke

import (
	core "jokes-bapak2-api/app/core/joke"
	"jokes-bapak2-api/app/core/schema"
	"jokes-bapak2-api/app/core/validator"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) UpdateJoke(c *fiber.Ctx) error {
	id := c.Params("id")
	// Check if the joke exists

	jokeExists, err := core.CheckJokeExists(d.DB, c.Context(), id)
	if err != nil {
		return err
	}

	if !jokeExists {
		return c.
			Status(fiber.StatusNotAcceptable).
			JSON(Error{
				Error: "specified joke id does not exists",
			})
	}

	body := new(schema.Joke)
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Check link validity
	valid, err := validator.CheckImageValidity(d.HTTP, body.Link)
	if err != nil {
		return err
	}

	if !valid {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(Error{
				Error: "URL provided is not a valid image",
			})
	}

	err = core.UpdateJoke(d.DB, c.Context(), body.Link, c.Locals("userID").(string))
	if err != nil {
		return err
	}

	err = core.SetAllJSONJoke(d.DB, c.Context(), d.Memory)
	if err != nil {
		return err
	}

	err = core.SetTotalJoke(d.DB, c.Context(), d.Memory)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(ResponseJoke{
			Message: "specified joke id has been updated",
			Link:    body.Link,
		})
}
