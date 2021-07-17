package models

type Error struct {
	Error string `json:"error"`
}

type ResponseJoke struct {
	Link    string `json:"link"`
	Message string `json:"message"`
}
