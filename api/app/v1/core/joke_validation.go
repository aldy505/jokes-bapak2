package core

import (
	"errors"
	"jokes-bapak2-api/app/v1/utils"
	"net/http"
	"strings"

	"github.com/gojek/heimdall/v7/httpclient"
)

var ValidContentType = []string{"image/jpeg", "image/png", "image/webp", "image/gif"}

// CheckImageValidity calls to the image host to check whether it's a valid image or not.
func CheckImageValidity(client *httpclient.Client, link string) (bool, error) {
	if strings.Contains(link, "https://") {
		// Perform HTTP call to link
		res, err := client.Get(link, http.Header{"User-Agent": []string{"JokesBapak2 API"}})
		if err != nil {
			return false, err
		}

		if res.StatusCode == 200 && utils.IsIn(ValidContentType, res.Header.Get("content-type")) {
			return true, nil
		}
		
		return false, nil
	}
	return false, errors.New("URL must use HTTPS protocol")
}