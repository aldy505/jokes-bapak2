package core

import (
	"bytes"
	"io"
	"io/ioutil"
	"jokes-bapak2-api/app/v1/utils"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/gojek/heimdall/v7/httpclient"
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

	var data ImageAPI
	err = ffjson.Unmarshal(responseBody, &data)
	if err != nil {
		return "", err
	}

	return data.Data.URL, nil
}
