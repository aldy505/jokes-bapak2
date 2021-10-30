package health

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dependencies struct {
	DB    *pgxpool.Pool
	Redis *redis.Client
}

func (d *Dependencies) Health(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(c.Context())
	if err != nil {
		return err
	}
	defer conn.Release()

	// Ping REDIS database
	err = d.Redis.Ping(c.Context()).Err()
	if err != nil {
		return c.
			Status(fiber.StatusServiceUnavailable).
			JSON(Error{
				Error: "REDIS: " + err.Error(),
			})
	}

	_, err = conn.Query(c.Context(), "SELECT \"id\" FROM \"jokesbapak2\" LIMIT 1")
	if err != nil {
		return c.
			Status(fiber.StatusServiceUnavailable).
			JSON(Error{
				Error: "POSTGRESQL: " + err.Error(),
			})
	}
	return c.SendStatus(fiber.StatusOK)
}
