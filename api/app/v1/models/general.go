package models

type Auth struct {
	ID       int    `json:"id" form:"id" db:"id"`
	Key      string `json:"key" form:"key" db:"key"`
	Token    string `json:"token" form:"token" db:"token"`
	LastUsed string `json:"last_used" form:"last_used" db:"last_used"`
}
