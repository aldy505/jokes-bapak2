package middleware

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	phccrypto "github.com/aldy505/phc-crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func RequireAuth(db *pgxpool.Pool, ctx *context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var auth Auth
		err := c.BodyParser(&auth)
		if err != nil {
			return err
		}

		// Check if key exists
		sql, args, err := psql.
			Select("token").
			From("administrators").
			Where(squirrel.Eq{"key": auth.Key}).
			ToSql()
		if err != nil {
			return err
		}

		var token string
		err = db.QueryRow(*ctx, sql, args...).Scan(&token)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return c.
					Status(fiber.StatusForbidden).
					JSON(Error{
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
			sql, args, err = psql.
				Update("administrators").
				Set("last_used", time.Now().UTC().Format(time.RFC3339)).
				ToSql()
			if err != nil {
				return err
			}

			_, err = db.Query(*ctx, sql, args...)
			if err != nil {
				return err
			}

			sql, args, err = psql.
				Select("id").
				From("administrators").
				Where(squirrel.Eq{"key": auth.Key}).
				ToSql()
			if err != nil {
				return err
			}

			var id int
			err = db.QueryRow(*ctx, sql, args...).Scan(&id)
			if err != nil {
				return err
			}
			c.Locals("userID", id)
			return c.Next()
		}

		return c.
			Status(fiber.StatusForbidden).
			JSON(Error{
				Error: "Invalid key",
			})
	}
}
