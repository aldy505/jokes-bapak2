package submit

import (
	"bytes"
	"net/url"
	"strconv"

	"github.com/aldy505/bob"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gofiber/fiber/v2"
)

func (d *Dependencies) GetSubmission(c *fiber.Ctx) error {
	query := new(SubmissionQuery)
	err := c.QueryParser(query)
	if err != nil {
		return err
	}

	var limit int
	var offset int
	var approved bool

	if query.Limit != "" {
		limit, err = strconv.Atoi(query.Limit)
		if err != nil {
			return err
		}
	}
	if query.Page != "" {
		page, err := strconv.Atoi(query.Page)
		if err != nil {
			return err
		}
		offset = (page - 1) * 20
	}

	if query.Approved != "" {
		approved, err = strconv.ParseBool(query.Approved)
		if err != nil {
			return err
		}
	}

	var status int

	if approved {
		status = 1
	} else {
		status = 0
	}

	var sql string
	var args []interface{}

	var sqlQuery *bytes.Buffer = &bytes.Buffer{}
	sqlQuery.WriteString("SELECT * FROM submission WHERE TRUE")

	if query.Author != "" {
		sqlQuery.WriteString(" AND author = ?")
		escapedAuthor, err := url.QueryUnescape(query.Author)
		if err != nil {
			return err
		}
		args = append(args, escapedAuthor)
	}

	if query.Approved != "" {
		sqlQuery.WriteString(" AND status = ?")
		args = append(args, status)
	}

	if limit > 0 {
		sqlQuery.WriteString(" LIMIT " + strconv.Itoa(limit))
	} else {
		sqlQuery.WriteString(" LIMIT 20")
	}

	if query.Page != "" {
		sqlQuery.WriteString(" OFFSET " + strconv.Itoa(offset))
	}

	sql = bob.ReplacePlaceholder(sqlQuery.String(), bob.Dollar)

	var submissions []Submission
	results, err := d.DB.Query(*d.Context, sql, args...)
	if err != nil {
		return err
	}

	defer results.Close()

	err = pgxscan.ScanAll(&submissions, results)
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
