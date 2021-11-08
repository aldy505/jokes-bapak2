package submit_test

import (
	"context"
	"jokes-bapak2-api/core/schema"
	"jokes-bapak2-api/core/submit"
	"testing"
)

func TestSubmitJoke(t *testing.T) {
	defer Flush()

	s, err := submit.SubmitJoke(db, context.Background(), schema.Submission{Author: "Test <example@test.com>"}, "https://example.net/img.png")
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if s.Link != "https://example.net/img.png" {
		t.Error("link is not correct, got:", s.Link)
	}
}
