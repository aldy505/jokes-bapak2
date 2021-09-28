package joke

import (
	"jokes-bapak2-api/app/core"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func (d *Dependencies) AddNewJoke(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(*d.Context)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(*d.Context)
	if err != nil {
		return err
	}
	defer tx.Rollback(*d.Context)

	var body core.Joke
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Check link validity
	valid, err := core.CheckImageValidity(d.HTTP, body.Link)
	if err != nil {
		return err
	}

	if !valid {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(Error{
				Error: "URL provided is not a valid image",
			})
	}
	// Validate if link already exists
	sql, args, err := d.Query.
		Select("link").
		From("jokesbapak2").
		Where(squirrel.Eq{"link": body.Link}).
		ToSql()
	if err != nil {
		return err
	}
	var validateLink string
	err = conn.QueryRow(*d.Context, sql, args...).Scan(&validateLink)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}
	
	if err == nil && validateLink != "" {
		return c.Status(fiber.StatusConflict).JSON(Error{
			Error: "Given link is already on the jokesbapak2 database",
		})
	}

	sql, args, err = d.Query.
		Insert("jokesbapak2").
		Columns("link", "creator").
		Values(body.Link, c.Locals("userID")).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(*d.Context, sql, args...)
	if err != nil {
		return err
	}

	err = tx.Commit(*d.Context)
	if err != nil {
		return err
	}

	err = core.SetAllJSONJoke(d.DB, d.Memory, d.Context)
	if err != nil {
		return err
	}
	err = core.SetTotalJoke(d.DB, d.Memory, d.Context)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(ResponseJoke{
			Link: body.Link,
		})
}
