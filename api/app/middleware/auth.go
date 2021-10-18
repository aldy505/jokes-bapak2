package middleware

import (
	"jokes-bapak2-api/app/core/administrator"

	phccrypto "github.com/aldy505/phc-crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func RequireAuth(db *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var auth Auth
		err := c.BodyParser(&auth)
		if err != nil {
			return err
		}

		token, err := administrator.CheckKeyExists(db, c.Context(), auth.Key)
		if err != nil {
			return err
		}

		if token == "" {
			return c.
				Status(fiber.StatusForbidden).
				JSON(Error{
					Error: "Invalid key",
				})
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
			id, err := administrator.GetUserID(db, c.Context(), auth.Key)
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
