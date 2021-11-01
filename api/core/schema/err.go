package schema

import "errors"

var ErrNotFound = errors.New("record not found")
var ErrEmpty = errors.New("record is empty")

type Error struct {
	Error string `json:"error"`
}
