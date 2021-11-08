package submit

import (
	"jokes-bapak2-api/core/schema"
	core "jokes-bapak2-api/core/submit"
	"jokes-bapak2-api/core/validator"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) SubmitJoke(c *fiber.Ctx) error {
	conn, err := d.DB.Acquire(c.Context())
	if err != nil {
		return err
	}
	defer conn.Release()

	var body schema.Submission
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Image and/or Link should not be empty
	if body.Image == "" && body.Link == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Error: "A link or an image should be supplied in a form of multipart/form-data",
		})
	}

	// Author should be supplied
	if body.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
			Error: "An author key consisting on the format \"yourname <youremail@mail>\" must be supplied",
		})
	} else {
		// Validate format
		valid := validator.ValidateAuthor(body.Author)
		if !valid {
			return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
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
			return c.Status(fiber.StatusBadRequest).JSON(schema.Error{
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
	validateLink, err := validator.SubmitLinkExists(d.DB, c.Context(), d.Query, link)
	if err != nil {
		return err
	}

	if validateLink {
		return c.Status(fiber.StatusConflict).JSON(schema.Error{
			Error: "Given link is already on the submission queue.",
		})
	}

	submission, err := core.SubmitJoke(d.DB, c.Context(), body, link)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(schema.ResponseSubmission{
			Message:    "Joke submitted. Please wait for a few days for admin to approve your submission.",
			Submission: submission,
			AuthorPage: "/submit?author=" + url.QueryEscape(body.Author),
		})
}
