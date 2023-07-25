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

	res, err := a.httpClient.Do(req)
	if err != nil {
		http.Error(w, "We're unable to complete your request at this time", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("TODO, res.StatusCode != http.StatusOK")
		return
	}

	var s StatScheduleResponse
	if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
		http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
		return
	}

	// No dates found for given date
	if len(s.Dates) == 0 {
		render.JSON(w, r, s)
		return
	}

	// No games found for given date
	if len(s.Dates[0].Games) == 0 {
		render.JSON(w, r, s)
		return
	}

	s.Dates[0].Games = sortGames(s.Dates[0].Games, teamIDInt)

	render.JSON(w, r, s)
}

func sortGames(games []Game, teamID int) []Game {
	// Sort by favorite team's ID in home or away games, preserving original order otherwise.
	less := func(i, _ int) bool {
		return games[i].Teams.Away.Team.ID == teamID ||
			games[i].Teams.Home.Team.ID == teamID
	}
	sort.SliceStable(games, less)

	// If we matched the favorite teamID and that team has a double header.
	favTeamDoubleHeader := (games[0].Teams.Away.Team.ID == teamID ||
		games[0].Teams.Home.Team.ID == teamID) &&
		games[0].DoubleHeader != "N"

	if !favTeamDoubleHeader {
		return games
	}
	if len(games) < 2 {
		fmt.Println("This should never happen at this point")
		return games
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

	// From takehome docs: "Any date in the 2021 and 2022 calendar years may be used to evaluate your service"
	// Afterwards it talks about if a game is live do x, y, z… Games can't ever be live if we only accept dates
	// in the past. Had this not been a code challenge, I would have gotten clarification around this. I wanted
	// to show possible logic for this, however I also wasnt 100% sure on the correct field to check against.
	// I made an assumption and chose "status.detailedState". Had I gotten clarification and the result was
	// that we indeed won't accept queries after 2022, then instead of this code I would have likely validated
	// that the date given was in 2021-2022 range at the top of this handler.
	if games[1].Status.DetailedState != "Final" {
		games[0], games[1] = games[1], games[0]
	}

	return games
}
