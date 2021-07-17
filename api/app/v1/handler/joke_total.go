package handler

import (
	"encoding/json"
	"jokes-bapak2-api/app/v1/core"
	"jokes-bapak2-api/app/v1/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func TotalJokes(c *fiber.Ctx) error {
	checkCache := core.CheckJokesCache(memory)

	if !checkCache {
		jokes, err := core.GetAllJSONJokes(db)
		if err != nil {
			return err
		}
		memory.Set("jokes", jokes, cache.NoExpiration)
	}

	jokes, found := memory.Get("jokes")
	if !found {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Error{
			Error: "no data found",
		})
	}

	var data []models.Joke
	err := json.Unmarshal(jokes.([]byte), &data)
	if err != nil {
		return err
	}

	dataLength := strconv.Itoa(len(data))
	return c.Status(fiber.StatusOK).JSON(models.ResponseJoke{
		Message: dataLength,
	})
}
