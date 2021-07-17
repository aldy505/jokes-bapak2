package middleware

import (
	"context"
	"time"

	"jokes-bapak2-api/app/v1/models"
	"jokes-bapak2-api/app/v1/platform/database"

	"github.com/Masterminds/squirrel"
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

		// Check if key exists
		sql, args, err := psql.Select("token").From("administrators").Where(squirrel.Eq{"key": auth.Key}).ToSql()
		if err != nil {
			return err
		}

		var token string
		err = db.QueryRow(context.Background(), sql, args...).Scan(&token)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return c.Status(fiber.StatusForbidden).JSON(models.Error{
					Error: "Invalid key",
				})
			}
			return err
		}

		crypto, err := phccrypto.Use(phccrypto.Argon2, phccrypto.Config{})
		if err != nil {
			return err
		}

		verify, err := crypto.Verify(token, auth.Token)
		if err != nil {
			return err
		}

		if verify {
			sql, args, err = psql.Update("administrator").Set("last_used", time.Now().UTC().Format(time.RFC3339)).ToSql()
			if err != nil {
				return err
			}

			_, err = db.Query(context.Background(), sql, args...)
			if err != nil {
				return err
			}

			sql, args, err = psql.Select("id").From("administrators").Where(squirrel.Eq{"key": auth.Key}).ToSql()
			if err != nil {
				return err
			}

			var id int
			err = db.QueryRow(context.Background(), sql, args...).Scan(&id)
			if err != nil {
				return err
			}
			c.Locals("userID", id)
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(models.Error{
			Error: "Invalid key",
		})
	}
}
