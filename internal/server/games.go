package server

import (
	"fmt"
	"sort"
)

func sortGames(games []Game, teamID int) []Game {
	// Sort by favorite team's ID in home or away games, preserving original order otherwise.
	less := func(i, _ int) bool {
		return games[i].Teams.Away.Team.ID == teamID ||
			games[i].Teams.Home.Team.ID == teamID
	}
	sort.SliceStable(games, less)

	// If we matched the favorite teamID and that team has a double header.
	favTeamDoubleHeader := (games[0].Teams.Away.Team.ID == teamID || games[0].Teams.Home.Team.ID == teamID) &&
		games[0].DoubleHeader != "N"

	if !favTeamDoubleHeader {
		return games
	}
	if len(games) < 2 {
		fmt.Println("TODO: This should never happen at this point as far as I can tell")
		return games
	}

	// If "single admission"/"traditional" doubleheader type and first game in slice is later by `startTimeTBD == true`
	game1LaterStartTime := games[0].DoubleHeader == "Y" && games[0].Status.StartTimeTBD
	// If "split admission" doubleheader type and the 1st game's `gameDate` is later than the 2nd games `gameDate`
	game1LaterGameDate := games[0].DoubleHeader == "S" && games[0].GameDate.After(games[1].GameDate)

	if game1LaterStartTime || game1LaterGameDate {
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
