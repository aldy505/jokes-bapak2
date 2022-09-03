package joke

import "time"

const JokesBapak2Bucket = "jokesbapak2"

type Joke struct {
	FileName    string
	ContentType string
	ModifiedAt  time.Time
}
