package joke

import (
	core "jokes-bapak2-api/core/joke"
	"net/http"
	"strconv"
)

func (d *Dependencies) TotalJokes(w http.ResponseWriter, r *http.Request) {
	total, err := core.GetTotalJoke(r.Context(), d.Bucket, d.Redis, d.Memory)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":` + strconv.Quote(err.Error()) + `}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":` + strconv.Itoa(total) + `}`))
}
