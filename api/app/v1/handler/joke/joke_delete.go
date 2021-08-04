package joke

import (
	"context"
	"strconv"

	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/handler"
	"jokes-bapak2-api/app/v1/models"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

func DeleteJoke(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Check if the joke exists
	sql, args, err := handler.Psql.Select("id").From("jokesbapak2").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	var jokeID int
	err = handler.Db.QueryRow(context.Background(), sql, args...).Scan(&jokeID)
	if err != nil {
		return err
	}

	if jokeID == id {
		sql, args, err = handler.Psql.Delete("jokesbapak2").Where(squirrel.Eq{"id": id}).ToSql()
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

		return c.Status(fiber.StatusOK).JSON(models.ResponseJoke{
			Message: "specified joke id has been deleted",
		})
	}
	return c.Status(fiber.StatusNotAcceptable).JSON(models.Error{
		Error: "specified joke id does not exists",
	})
}
