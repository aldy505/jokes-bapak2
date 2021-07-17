package models

type Error struct {
	Error string `json:"error"`
}

type ResponseJoke struct {
	Link    string `json:"link,omitempty"`
	Message string `json:"message,omitempty"`
}
