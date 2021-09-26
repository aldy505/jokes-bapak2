package joke

type ResponseJoke struct {
	Link    string `json:"link,omitempty"`
	Message string `json:"message,omitempty"`
}

type Today struct {
	Date        string `redis:"today:date"`
	Image       string `redis:"today:image"`
	ContentType string `redis:"today:contentType"`
}

type Error struct {
	Error string `json:"error"`
}
