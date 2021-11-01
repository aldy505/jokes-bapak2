package joke

import (
	core "jokes-bapak2-api/core/joke"
	"jokes-bapak2-api/core/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) DeleteJoke(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	validate, err := validator.JokeIDExists(d.DB, c.Context(), id)
	if err != nil {
		return err
	}

	if validate {
		return c.
			Status(fiber.StatusNotAcceptable).
			JSON(Error{
				Error: "specified joke id does not exists",
			})
	}

	err = core.DeleteSingleJoke(d.DB, c.Context(), id)
	if err != nil {
		return err
	}

	err = core.SetAllJSONJoke(d.DB, c.Context(), d.Memory)
	if err != nil {
		return err
	}

	err = core.SetTotalJoke(d.DB, c.Context(), d.Memory)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(ResponseJoke{
			Message: "specified joke id has been deleted",
		})

}
