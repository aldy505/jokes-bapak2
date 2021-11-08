package joke

import (
	core "jokes-bapak2-api/core/joke"
	"jokes-bapak2-api/core/schema"
	"jokes-bapak2-api/core/validator"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) AddNewJoke(c *fiber.Ctx) error {
	var body schema.Joke
	err := c.BodyParser(&body)
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

	validateLink, err := validator.JokeLinkExists(d.DB, c.Context(), body.Link)
	if err != nil {
		return err
	}

	if !validateLink {
		return c.Status(fiber.StatusConflict).JSON(Error{
			Error: "Given link is already on the jokesbapak2 database",
		})
	}

	err = core.InsertJokeIntoDB(
		d.DB,
		c.Context(),
		schema.Joke{
			Link:    body.Link,
			Creator: c.Locals("userID").(int),
		},
	)
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
		Status(fiber.StatusCreated).
		JSON(ResponseJoke{
			Link: body.Link,
		})
}
