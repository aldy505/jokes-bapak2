package joke

import "time"

// JokesBapak2Bucket defines the bucket that the jokes resides in.
const JokesBapak2Bucket = "jokesbapak2"

// Joke provides a simple struct that points
// to the information of the joke.
type Joke struct {
	FileName    string
	ContentType string
	ModifiedAt  time.Time
}
