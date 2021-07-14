package models

type RequestJokePost struct {
	Link string `json:"link" form:"link"`
}

type RequestAuth struct {
	Key   string `json:"key" form:"key"`
	Token string `json:"token" form:"token"`
}

type Today struct {
	Date        string `redis:"today:date"`
	Image       string `redis:"today:image"`
	ContentType string `redis:"today:contentType"`
}
