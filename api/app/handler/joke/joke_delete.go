package joke

import (
	"jokes-bapak2-api/app/core"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) DeleteJoke(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(*d.Context)
	if err != nil {
		return err
	}
	defer conn.Release()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Check if the joke exists
	sql, args, err := d.Query.
		Select("id").
		From("jokesbapak2").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	var jokeID int
	err = conn.QueryRow(*d.Context, sql, args...).Scan(&jokeID)
	if err != nil {
		return err
	}

	if jokeID == id {
		sql, args, err = d.Query.
			Delete("jokesbapak2").
			Where(squirrel.Eq{"id": id}).
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
				Message: "specified joke id has been deleted",
			})
	}
	return c.
		Status(fiber.StatusNotAcceptable).
		JSON(Error{
			Error: "specified joke id does not exists",
		})
}
