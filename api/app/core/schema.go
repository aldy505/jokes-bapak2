package core

import "errors"

type Joke struct {
	ID      int    `json:"id" form:"id" db:"id"`
	Link    string `json:"link" form:"link" db:"link"`
	Creator int    `json:"creator" form:"creator" db:"creator"`
}

type ImageAPI struct {
	Data    ImageAPIData `json:"data"`
	Success bool         `json:"success"`
	Status  int          `json:"status"`
}

type ImageAPIData struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URLViewer  string `json:"url_viewer"`
	URL        string `json:"url"`
	DisplayURL string `json:"display_url"`
}

var ErrNotFound = errors.New("record not found")
var ErrEmpty = errors.New("record is empty")
