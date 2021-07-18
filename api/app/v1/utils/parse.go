package utils

import (
	"strconv"

	"github.com/pquerna/ffjson/ffjson"
)

// ParseToFormBody converts a body to form data type
func ParseToFormBody(body map[string]interface{}) ([]byte, error) {
	var form string
	for key, value := range body {
		form += key + "="
		switch v := value.(type) {
		case string:
			form += v
		case int:
			form += strconv.Itoa(v)
		case bool:
			form += strconv.FormatBool(v)
		}
		form += "&"
	}
	return []byte(form), nil
}

// ParseToJSONBody converts a body to json data type
func ParseToJSONBody(body map[string]interface{}) ([]byte, error) {
	b, err := ffjson.Marshal(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
