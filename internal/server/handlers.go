package server

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func (a *API) ping(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) }

func (a *API) getSchedule(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	teamID := r.URL.Query().Get("teamID")

	layout := "2006-01-02"
	if _, err := time.Parse(layout, date); err != nil {
		a.log.Warn(err.Error(), zap.String("date", date))
		http.Error(w, "Invalid date format, expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	teamIDInt, err := strconv.Atoi(teamID)
	if err != nil {
		a.log.Warn(err.Error(), zap.String("teamID", teamID))
		http.Error(w, "Invalid teamID format, expected a number", http.StatusBadRequest)
		return
	}

	params := url.Values{}
	params.Add("date", date)
	params.Add("sportId", "1")
	params.Add("language", "en")
	mlbStatsURL := mlbStatsBaseURL + "/api/v1/schedule?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, mlbStatsURL, nil)
	if err != nil {
		a.log.Error(err.Error(), zap.String("method", http.MethodGet), zap.String("URL", mlbStatsURL))
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Accept", "application/json")

	res, err := a.httpClient.Do(req)
	if err != nil {
		a.log.Error(err.Error())
		http.Error(w, "We're unable to complete your request at this time", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	// I would need some more context around what exactly we want to do with other types of
	// return codes from the MLB API. Going with internal server error for the time being
	// when we don't simply get an OK (200).
	if res.StatusCode != http.StatusOK {
		a.log.Warn("Return status code not OK", zap.Int("statusCode", res.StatusCode))
		http.Error(w, "We're unable to complete your request at this time", http.StatusInternalServerError)
		return
	}

	var s StatScheduleResponse
	if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
		a.log.Error(err.Error())
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	if len(s.Dates) == 0 || len(s.Dates[0].Games) == 0 {
		render.JSON(w, r, s)
		return
	}

	s.Dates[0].Games = sortGames(s.Dates[0].Games, teamIDInt)

	render.JSON(w, r, s)
}
