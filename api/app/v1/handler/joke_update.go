package handler

import (
	"context"

	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/models"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func UpdateJoke(c *fiber.Ctx) error {
	id := c.Params("id")
	// Check if the joke exists
	sql, args, err := psql.Select("id").From("jokesbapak2").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	var jokeID string
	err = db.QueryRow(context.Background(), sql, args...).Scan(&jokeID)
	if err != nil && err != models.ErrNoRows {
		return err
	}

	if jokeID == id {
		body := new(models.Joke)
		err = c.BodyParser(&body)
		if err != nil {
			return err
		}

		sql, args, err = psql.Update("jokesbapak2").Set("link", body.Link).Set("creator", c.Locals("userID")).ToSql()
		if err != nil {
			return err
		}

		_, err := db.Query(context.Background(), sql, args...)
		if err != nil {
			return err
		}

		jokes, err := core.GetAllJSONJokes(db)
		if err != nil {
			return err
		}

		err = memory.Set("jokes", jokes)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(models.ResponseJoke{
			Message: "specified joke id has been updated",
			Link:    body.Link,
		})
	}

	return c.Status(fiber.StatusNotAcceptable).JSON(models.Error{
		Error: "specified joke id does not exists",
	})
}
