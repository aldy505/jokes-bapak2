package joke

import (
	core "jokes-bapak2-api/core/joke"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// TodayJoke provides http handler for today's joke
func (d *Dependencies) TodayJoke(w http.ResponseWriter, r *http.Request) {
	joke, contentType, err := core.GetTodaysJoke(r.Context(), d.Bucket, d.Redis, d.Memory)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":` + strconv.Quote(err.Error()) + `}`))
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(joke)
}

// SingleJoke provides http handler for acquiring random single joke
func (d *Dependencies) SingleJoke(w http.ResponseWriter, r *http.Request) {
	joke, contentType, err := core.GetRandomJoke(r.Context(), d.Bucket, d.Redis, d.Memory)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":` + strconv.Quote(err.Error()) + `}`))
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(joke)

}

// JokeByID provides http handler for acquiring a joke by ID
func (d *Dependencies) JokeByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParamFromCtx(r.Context(), "id")

	// Parse id to int
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":` + strconv.Quote(err.Error()) + `}`))
		return
	}

	joke, contentType, err := core.GetJokeByID(r.Context(), d.Bucket, d.Redis, d.Memory, parsedId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":` + strconv.Quote(err.Error()) + `}`))
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(joke)
}
