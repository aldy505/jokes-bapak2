package joke

import (
	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/handler"
	"jokes-bapak2-api/app/v1/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func TotalJokes(c *fiber.Ctx) error {
	checkTotal, err := core.CheckTotalJokesCache(handler.Memory)
	if err != nil {
		return err
	}

	if !checkTotal {
		err = core.SetTotalJoke(handler.Db, handler.Memory)
		if err != nil {
			return err
		}
	}

	total, err := handler.Memory.Get("total")

	if err != nil {
		if err.Error() == "Entry not found" {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(models.Error{
					Error: "no data found",
				})
		}
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(models.ResponseJoke{
			Message: strconv.Itoa(int(total[0])),
		})
}
