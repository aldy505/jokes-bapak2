package joke

import (
	"errors"
	core "jokes-bapak2-api/core/joke"
	"strconv"

	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) TotalJokes(c *fiber.Ctx) error {
	checkTotal, err := core.CheckTotalJokesCache(d.Memory)
	if err != nil {
		return err
	}

	if !checkTotal {
		err = core.SetTotalJoke(d.DB, c.Context(), d.Memory)
		if err != nil {
			return err
		}
	}

	total, err := d.Memory.Get("total")

	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return c.
				Status(fiber.StatusInternalServerError).
				JSON(Error{
					Error: "no data found",
				})
		}
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(ResponseJoke{
			Message: strconv.Itoa(int(total[0])),
		})
}
