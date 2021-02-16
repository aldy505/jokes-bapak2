package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	limiter "github.com/ulule/limiter/v3"
)

func main() {
	var connStr string = "postgres://"
	db, _ := sql.Open("postgres", connStr)

	app := fiber.New()
	app.Use(cors.Default())
	app.Use(limiter.New().handler(app))

	app.Get("/:any", func(c *fiber.Ctx) error {
		return c.SendFile("")
	})
	app.Listen(":3000")
}
