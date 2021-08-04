package joke

import (
	"context"

	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/handler"
	"jokes-bapak2-api/app/v1/models"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func UpdateJoke(c *fiber.Ctx) error {
	id := c.Params("id")
	// Check if the joke exists
	sql, args, err := handler.Psql.
		Select("id").
		From("jokesbapak2").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	var jokeID string
	err = handler.Db.QueryRow(context.Background(), sql, args...).Scan(&jokeID)
	if err != nil && err != models.ErrNoRows {
		return err
	}

	if jokeID == id {
		body := new(models.Joke)
		err = c.BodyParser(&body)
		if err != nil {
			return err
		}

		// Check link validity
		valid, err := core.CheckImageValidity(handler.Client, body.Link)
		if err != nil {
			return err
		}

		if !valid {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(models.Error{
					Error: "URL provided is not a valid image",
				})
		}

		sql, args, err = handler.Psql.
			Update("jokesbapak2").
			Set("link", body.Link).
			Set("creator", c.Locals("userID")).
			ToSql()
		if err != nil {
			return err
		}

		r, err := handler.Db.Query(context.Background(), sql, args...)
		if err != nil {
			return err
		}

		defer r.Close()

		err = core.SetAllJSONJoke(handler.Db, handler.Memory)
		if err != nil {
			return err
		}
		err = core.SetTotalJoke(handler.Db, handler.Memory)
		if err != nil {
			return err
		}

		return c.
			Status(fiber.StatusOK).
			JSON(models.ResponseJoke{
				Message: "specified joke id has been updated",
				Link:    body.Link,
			})
	}

	return c.
		Status(fiber.StatusNotAcceptable).
		JSON(models.Error{
			Error: "specified joke id does not exists",
		})
}
