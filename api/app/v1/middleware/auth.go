package middleware

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/models"
	"github.com/aldy505/jokes-bapak2-api/api/app/v1/platform/database"
	phccrypto "github.com/aldy505/phc-crypto"
	"github.com/gofiber/fiber/v2"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
var db = database.New()

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var auth models.Auth
		err := c.BodyParser(&auth)
		if err != nil {
			return err
		}

		// Check if token exists
		sql, args, err := psql.Select("token").From("authorization").Where("token", auth.Token).ToSql()
		if err != nil {
			return err
		}
		var token string
		err = db.QueryRow(context.Background(), sql, args...).Scan(&token)
		if err != nil {
			return err
		}

		if token == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		sql, args, err = psql.Select("key").From("authorization").Where("token", token).ToSql()
		if err != nil {
			return err
		}
		var key string
		err = db.QueryRow(context.Background(), sql, args...).Scan(&key)
		if err != nil {
			return err
		}

		crypto, err := phccrypto.Use(phccrypto.Argon2, phccrypto.Config{})
		if err != nil {
			return err
		}

		verify, err := crypto.Verify(key, auth.Key)
		if err != nil {
			return err
		}

		if verify {
			c.Next()
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Invalid key",
		})
	}
}
