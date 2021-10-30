package schema

type Joke struct {
	ID      int    `json:"id" form:"id" db:"id"`
	Link    string `json:"link" form:"link" db:"link"`
	Creator int    `json:"creator" form:"creator" db:"creator"`
}
