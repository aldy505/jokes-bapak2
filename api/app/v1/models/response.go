package models

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseJoke struct {
	Link    string `json:"link"`
	Message string `json:"message"`
}
