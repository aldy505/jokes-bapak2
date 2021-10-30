package utils

import (
	"bytes"
	"net/http"
)

type ContentType int

const (
	JSON ContentType = iota
	Form
)

type RequestConfig struct {
	URL         string
	Method      string
	Headers     map[string]interface{}
	Body        map[string]interface{}
	ContentType ContentType
}

// Request is a simple wrapper around http.NewRequest
func Request(config RequestConfig) (response *http.Response, err error) {
	client := &http.Client{}

	var body []byte
	if config.ContentType == JSON {
		parsed, err := ParseToJSONBody(config.Body)
		if err != nil {
			return &http.Response{}, err
		}

		body = parsed
	} else if config.ContentType == Form {
		parsed, err := ParseToFormBody(config.Body)
		if err != nil {
			return &http.Response{}, err
		}

		body = parsed
	}

	request, err := http.NewRequest(config.Method, config.URL, bytes.NewReader(body))
	if err != nil {
		return
	}

	response, err = client.Do(request)
	if err != nil {
		return
	}

	return
}
