package joke

import (
	"jokes-bapak2-api/app/core"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func (d *Dependencies) UpdateJoke(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(*d.Context)
	if err != nil {
		return err
	}
	defer conn.Release()

	id := c.Params("id")
	// Check if the joke exists
	sql, args, err := d.Query.
		Select("id").
		From("jokesbapak2").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	var jokeID string
	err = conn.QueryRow(*d.Context, sql, args...).Scan(&jokeID)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if jokeID == id {
		body := new(core.Joke)
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

		sql, args, err = d.Query.
			Update("jokesbapak2").
			Set("link", body.Link).
			Set("creator", c.Locals("userID")).
			ToSql()
		if err != nil {
			return err
		}

		r, err := conn.Query(*d.Context, sql, args...)
		if err != nil {
			return err
		}
		defer r.Close()

		err = core.SetAllJSONJoke(d.DB, d.Memory, d.Context)
		if err != nil {
			return err
		}
		err = core.SetTotalJoke(d.DB, d.Memory, d.Context)
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

	return c.
		Status(fiber.StatusNotAcceptable).
		JSON(Error{
			Error: "specified joke id does not exists",
		})
}
