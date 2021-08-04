package models

import "errors"

var ErrNoRows = errors.New("no rows in result set")
var ErrConnDone = errors.New("connection is already closed")
var ErrTxDone = errors.New("transaction has already been committed or rolled back")

var ErrNotFound = errors.New("record not found")
var ErrEmpty = errors.New("record is empty")

type Error struct {
	Error string `json:"error"`
}
