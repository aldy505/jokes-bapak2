package joke

import (
	"jokes-bapak2-api/app/core"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) AddNewJoke(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(*d.Context)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(*d.Context)
	if err != nil {
		return err
	}
	defer tx.Rollback(*d.Context)

	var body core.Joke
	err = c.BodyParser(&body)
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
	_, err = tx.Exec(*d.Context, sql, args...)
	if err != nil {
		return err
	}

	err = tx.Commit(*d.Context)
	if err != nil {
		return err
	}

	err = core.SetAllJSONJoke(d.DB, d.Memory, d.Context)
	if err != nil {
		return err
	}
	err = core.SetTotalJoke(d.DB, d.Memory, d.Context)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(ResponseJoke{
			Link: body.Link,
		})
}
