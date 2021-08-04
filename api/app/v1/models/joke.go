package models

type Joke struct {
	ID      int    `json:"id" form:"id" db:"id"`
	Link    string `json:"link" form:"link" db:"link"`
	Creator int    `json:"creator" form:"creator" db:"creator"`
}

type Today struct {
	Date        string `redis:"today:date"`
	Image       string `redis:"today:image"`
	ContentType string `redis:"today:contentType"`
}

type ResponseJoke struct {
	Link    string `json:"link,omitempty"`
	Message string `json:"message,omitempty"`
}
