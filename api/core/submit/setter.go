package submit

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"jokes-bapak2-api/core/schema"
	"jokes-bapak2-api/utils"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pquerna/ffjson/ffjson"
)

// UploadImage process the image from the user to be uploaded to the cloud storage.
// Returns the image URL.
func UploadImage(client *httpclient.Client, image io.Reader) (string, error) {
	hostURL := os.Getenv("IMAGE_API_URL")
	fileName, err := utils.RandomString(10)
	if err != nil {
		return "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("image")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(fw, image)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	headers := http.Header{
		"Content-Type": []string{writer.FormDataContentType()},
		"User-Agent":   []string{"JokesBapak2 API"},
		"Accept":       []string{"application/json"},
	}

	requestURL, err := url.Parse(hostURL)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("key", os.Getenv("IMAGE_API_KEY"))
	params.Add("name", fileName)

	requestURL.RawQuery = params.Encode()

	res, err := client.Post(requestURL.String(), bytes.NewReader(body.Bytes()), headers)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var data schema.ImageAPI
	err = ffjson.Unmarshal(responseBody, &data)
	if err != nil {
		return "", err
	}

	return data.Data.URL, nil
}

func SubmitJoke(db *pgxpool.Pool, ctx context.Context, s schema.Submission, link string) (schema.Submission, error) {
	var query = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	conn, err := db.Acquire(ctx)
	if err != nil {
		return schema.Submission{}, err
	}
	defer conn.Release()

	now := time.Now().UTC().Format(time.RFC3339)

	sql, args, err := query.
		Insert("submission").
		Columns("link", "created_at", "author").
		Values(link, now, s.Author).
		Suffix("RETURNING id,created_at,link,author,status").
		ToSql()
	if err != nil {
		return schema.Submission{}, err
	}

	var submission schema.Submission
	result, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return schema.Submission{}, err
	}
	defer result.Close()

	err = pgxscan.ScanOne(&submission, result)
	if err != nil {
		return schema.Submission{}, err
	}

	return submission, nil
}
