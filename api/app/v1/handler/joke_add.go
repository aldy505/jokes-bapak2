package handler

import (
	"context"

	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/models"

	"github.com/gofiber/fiber/v2"
)

func AddNewJoke(c *fiber.Ctx) error {
	var body models.Joke
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Check link validity
	valid, err := core.CheckImageValidity(client, body.Link)
	if err != nil {
		return err
	}

	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Error: "URL provided is not a valid image",
		})
	}

	sql, args, err := psql.Insert("jokesbapak2").Columns("link", "creator").Values(body.Link, c.Locals("userID")).ToSql()
	if err != nil {
		return err
	}

	// TODO: Implement solution if the link provided already exists.
	_, err = db.Query(context.Background(), sql, args...)
	if err != nil {
		return err
	}

	err = core.SetAllJSONJoke(db, memory)
	if err != nil {
		return err
	}
	err = core.SetTotalJoke(db, memory)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(models.ResponseJoke{
		Link: body.Link,
	})
}
