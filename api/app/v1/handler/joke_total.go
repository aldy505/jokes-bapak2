package handler

import (
	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/ffjson/ffjson"
)

func TotalJokes(c *fiber.Ctx) error {
	checkCache, err := core.CheckJokesCache(memory)
	if err != nil {
		return err
	}

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(db)
		if err != nil {
			return err
		}
		err = memory.Set("jokes", jokes)
		if err != nil {
			return err
		}
	}

	jokes, err := memory.Get("jokes")
	if err != nil {
		if err.Error() == "Entry not found" {
			return c.Status(fiber.StatusInternalServerError).JSON(models.Error{
				Error: "no data found",
			})
		}
		return err
	}

	var data []models.Joke
	err = ffjson.Unmarshal(jokes, &data)
	if err != nil {
		return err
	}

	dataLength := strconv.Itoa(len(data))
	return c.Status(fiber.StatusOK).JSON(models.ResponseJoke{
		Message: dataLength,
	})
}
