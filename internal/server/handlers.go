package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (a *API) getSchedule(w http.ResponseWriter, r *http.Request) {
	// Query params: teamID, date

	req, err := http.NewRequest("GET", "https://statsapi.mlb.com/api/v1/schedule?date=2021-09-19&sportId=1&language=en", nil)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Accept", "application/json")

	res, err := a.httpClient.Do(req)
	if err != nil {
		// Torn between 500 and 502. It isn't exactly a proxy, as we do some
		// additional logic/processing, but 502 seems to fit best here.
		http.Error(w, "Error retrieving data from source", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	// if res.StatusCode != http.StatusOK {}

	var s StatScheduleResponse
	if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, fmt.Sprintf(`{ "testStringData": %s }`, s.Dates[0].Games[0].Teams.Away.Team.Name))
}
