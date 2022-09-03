package schema

type Joke struct {
	ID      int    `json:"id"`
	Link    string `json:"link"`
	Creator int    `json:"creator"`
}
