package health

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dependencies struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
}

func (d *Dependencies) Health(c *fiber.Ctx) error {
	// Ping REDIS database
	err := d.Redis.Ping(context.Background()).Err()
	if err != nil {
		return c.
			Status(fiber.StatusServiceUnavailable).
			JSON(Error{
				Error: "REDIS: " + err.Error(),
			})
	}

	_, err = d.DB.Query(context.Background(), "SELECT \"id\" FROM \"jokesbapak2\" LIMIT 1")
	if err != nil {
		return c.
			Status(fiber.StatusServiceUnavailable).
			JSON(Error{
				Error: "POSTGRESQL: " + err.Error(),
			})
	}
	return c.SendStatus(fiber.StatusOK)
}
