package joke

import (
	"context"

	"jokes-bapak2-api/app/v1/core"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) AddNewJoke(c *fiber.Ctx) error {
	var body core.Joke
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Check link validity
	valid, err := core.CheckImageValidity(d.HTTP, body.Link)
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

	sql, args, err := d.Query.
		Insert("jokesbapak2").
		Columns("link", "creator").
		Values(body.Link, c.Locals("userID")).
		ToSql()
	if err != nil {
		return err
	}

	// TODO: Implement solution if the link provided already exists.
	r, err := d.DB.Query(context.Background(), sql, args...)
	if err != nil {
		return err
	}

	defer r.Close()

	err = core.SetAllJSONJoke(d.DB, d.Memory)
	if err != nil {
		return err
	}
	err = core.SetTotalJoke(d.DB, d.Memory)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(ResponseJoke{
			Link: body.Link,
		})
}
