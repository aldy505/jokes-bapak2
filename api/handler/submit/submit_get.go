package submit

import (
	"jokes-bapak2-api/core/schema"
	core "jokes-bapak2-api/core/submit"

	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) GetSubmission(c *fiber.Ctx) error {
	query := new(schema.SubmissionQuery)
	err := c.QueryParser(query)
	if err != nil {
		return err
	}

	submissions, err := core.GetSubmittedItems(d.DB, c.Context(), *query)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"count": len(submissions),
			"jokes": submissions,
		})
}
