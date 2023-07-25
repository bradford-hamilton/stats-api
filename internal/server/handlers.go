package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/go-chi/render"
)

func (a *API) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (a *API) getSchedule(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	teamID := r.URL.Query().Get("teamID")

	// Validate query params
	layout := "2006-01-02"
	if _, err := time.Parse(layout, date); err != nil {
		http.Error(w, "Invalid date format, expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	teamIDInt, err := strconv.Atoi(teamID)
	if err != nil {
		http.Error(w, "Invalid teamID format, expected a number", http.StatusBadRequest)
		return
	}

	// Build http request
	params := url.Values{}
	params.Add("date", date)
	params.Add("sportId", "1")
	params.Add("language", "en")
	mlbStatsURL := mlbStatsBaseURL + "/api/v1/schedule?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, mlbStatsURL, nil)
	if err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Accept", "application/json")

	// Make call to mlb stats API
	res, err := a.httpClient.Do(req)
	if err != nil {
		// Torn between 500 and 502. It isn't exactly a proxy, as we do some
		// additional logic/processing, but 502 seems to fit best here.
		http.Error(w, "We're unable to complete your request at this time", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	// Without docs or requirements around other responses,
	// I'm just going to make sure we got a 200
	if res.StatusCode != http.StatusOK {
		fmt.Println("TODO, res.StatusCode != http.StatusOK")
		return
	}

	var s StatScheduleResponse
	if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	// Filter results based on requirements using teamID.
	if len(s.Dates) == 0 {
		// Something went wrong/hard to define without MLB API docs.
		fmt.Println("TODO, len(s.Dates) == 0")
		fmt.Println("TODO No games this day?")
		return
	}

	games := s.Dates[0].Games
	if len(games) == 0 {
		// Something went wrong/hard to define without MLB API docs.
		fmt.Println("TODO, len(games) == 0")
		return
	}

	// Sort by favorite team's ID in home or away games, preserving original order.
	less := func(i, _ int) bool {
		return games[i].Teams.Away.Team.ID == teamIDInt ||
			games[i].Teams.Home.Team.ID == teamIDInt
	}
	sort.SliceStable(games, less)

	// // If we matched the favorite teamID and that team has a double header
	favTeamDoubleHeader := (games[0].Teams.Away.Team.ID == teamIDInt ||
		games[0].Teams.Home.Team.ID == teamIDInt) &&
		games[0].DoubleHeader != "N"

	if favTeamDoubleHeader {
		if len(games) < 2 {
			// Something odd happened, but don't crash TODO
			fmt.Printf("len(games) == %d\n", len(games))
		}

		if games[0].DoubleHeader == "Y" && games[0].Status.StartTimeTBD {
			// If "single admission"/"traditional" doubleheader type,
			// sort games 1 and 2 chronologically using "startTimeTBD"
			games[0], games[1] = games[1], games[0]
		} else if games[0].DoubleHeader == "S" && (games[0].GameDate.After(games[1].GameDate)) {
			// If "split admission" doubleheader type, sort games
			// 1 and 2 chronologically using "gameDate"
			games[0], games[1] = games[1], games[0]
		}

		// TODO After those sorts, while still under if favTeamDoubleHeader if the second game is "live", move that up to games[0]

	}

	render.JSON(w, r, s)
}
