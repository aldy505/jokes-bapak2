package models

type Joke struct {
	ID      int    `json:"id" form:"id" db:"id"`
	Link    string `json:"link" form:"link" db:"link"`
	Creator int    `json:"creator" form:"creator" db:"creator"`
}

type Auth struct {
	ID       int    `json:"id" form:"id" db:"id"`
	Key      string `json:"key" form:"key" db:"key"`
	Token    string `json:"token" form:"token" db:"token"`
	LastUsed string `json:"last_used" form:"last_used" db:"last_used"`
}

type Today struct {
	Date        string `redis:"today:date"`
	Image       string `redis:"today:image"`
	ContentType string `redis:"today:contentType"`
}
