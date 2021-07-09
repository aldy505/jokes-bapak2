package models

type RequestJokePost struct {
	Key  string `json:"string"`
	Link string `json:"link"`
}

type RequestAuth struct {
	Key   string `json:"key"`
	Token string `json:"token"`
}
