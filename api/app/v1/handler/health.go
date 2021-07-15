package handler

import (
	"context"
	"jokes-bapak2-api/app/v1/models"

	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	// Ping REDIS database
	err := redis.Ping(context.Background()).Err()
	if err != nil {
		return c.
			Status(fiber.StatusServiceUnavailable).
			JSON(models.ResponseError{
				Error: "REDIS: " + err.Error(),
			})
	}

	_, err = db.Query(context.Background(), "SELECT \"id\" FROM \"jokesbapak2\" LIMIT 1")
	if err != nil {
		return c.
			Status(fiber.StatusServiceUnavailable).
			JSON(models.ResponseError{
				Error: "POSTGRESQL: " + err.Error(),
			})
	}
	return c.SendStatus(fiber.StatusOK)
}
