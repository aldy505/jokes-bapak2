package submit

import (
	"context"
	core "jokes-bapak2-api/app/core/submit"
	"jokes-bapak2-api/app/core/validator"
	"net/url"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func (d *Dependencies) SubmitJoke(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(c.Context())
	if err != nil {
		return err
	}
	defer conn.Release()

	var body Submission
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Image and/or Link should not be empty
	if body.Image == "" && body.Link == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Error{
			Error: "A link or an image should be supplied in a form of multipart/form-data",
		})
	}

	// Author should be supplied
	if body.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(Error{
			Error: "An author key consisting on the format \"yourname <youremail@mail>\" must be supplied",
		})
	} else {
		// Validate format
		valid := validator.ValidateAuthor(body.Author)
		if !valid {
			return c.Status(fiber.StatusBadRequest).JSON(Error{
				Error: "Please stick to the format of \"yourname <youremail@mail>\" and within 200 characters",
			})
		}
	}

	var link string

	// Check link validity if link was provided
	if body.Link != "" {
		valid, err := validator.CheckImageValidity(d.HTTP, body.Link)
		if err != nil {
			return err
		}
		if !valid {
			return c.Status(fiber.StatusBadRequest).JSON(Error{
				Error: "URL provided is not a valid image",
			})
		}

		link = body.Link
	}

	// If image was provided
	if body.Image != "" {
		image := strings.NewReader(body.Image)

		link, err = core.UploadImage(d.HTTP, image)
		if err != nil {
			return err
		}
	}

	// Validate if link already exists
	validateLink, err := validateIfLinkExists(d.DB, c.Context(), d.Query, link)
	if err != nil {
		return err
	}

	if validateLink {
		return c.Status(fiber.StatusConflict).JSON(Error{
			Error: "Given link is already on the submission queue.",
		})
	}

	now := time.Now().UTC().Format(time.RFC3339)

	sql, args, err := d.Query.
		Insert("submission").
		Columns("link", "created_at", "author").
		Values(link, now, body.Author).
		Suffix("RETURNING id,created_at,link,author,status").
		ToSql()
	if err != nil {
		return err
	}

	var submission []Submission
	result, err := conn.Query(c.Context(), sql, args...)
	if err != nil {
		return err
	}
	defer result.Close()

	err = pgxscan.ScanAll(&submission, result)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(ResponseSubmission{
			Message:    "Joke submitted. Please wait for a few days for admin to approve your submission.",
			Submission: submission[0],
			AuthorPage: "/submit?author=" + url.QueryEscape(body.Author),
		})
}

func validateIfLinkExists(db *pgxpool.Pool, ctx context.Context, query squirrel.StatementBuilderType, link string) (bool, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Release()

	sql, args, err := query.
		Select("link").
		From("submission").
		Where(squirrel.Eq{"link": link}).
		ToSql()
	if err != nil {
		return false, err
	}

	var validateLink string
	err = conn.QueryRow(context.Background(), sql, args...).Scan(&validateLink)
	if err != nil && err != pgx.ErrNoRows {
		return false, err
	}

	if err == nil && validateLink != "" {
		return true, nil
	}

	return false, nil
}
