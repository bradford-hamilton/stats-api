package server

import "time"

// StatScheduleResponse represents the structure of the JSON returned from http
// calls to the mlb stats API (https://statsapi.mlb.com/api/v1/schedule)
type StatScheduleResponse struct {
	Copyright            string `json:"copyright"`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Dates                []struct {
		Date                 string        `json:"date"`
		TotalItems           int           `json:"totalItems"`
		TotalEvents          int           `json:"totalEvents"`
		TotalGames           int           `json:"totalGames"`
		TotalGamesInProgress int           `json:"totalGamesInProgress"`
		Games                []Game        `json:"games"`
		Events               []interface{} `json:"events"`
	} `json:"dates"`
}

// Game represents a single game from within a StatScheduleResponse
type Game struct {
	GamePk       int       `json:"gamePk"`
	Link         string    `json:"link"`
	GameType     string    `json:"gameType"`
	Season       string    `json:"season"`
	GameDate     time.Time `json:"gameDate"`
	OfficialDate string    `json:"officialDate"`
	Status       struct {
		AbstractGameState string `json:"abstractGameState"`
		CodedGameState    string `json:"codedGameState"`
		DetailedState     string `json:"detailedState"`
		StatusCode        string `json:"statusCode"`
		StartTimeTBD      bool   `json:"startTimeTBD"`
		AbstractGameCode  string `json:"abstractGameCode"`
	} `json:"status"`
	Teams struct {
		Away struct {
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Pct    string `json:"pct"`
			} `json:"leagueRecord"`
			Score int `json:"score"`
			Team  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			IsWinner     bool `json:"isWinner"`
			SplitSquad   bool `json:"splitSquad"`
			SeriesNumber int  `json:"seriesNumber"`
		} `json:"away"`
		Home struct {
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Pct    string `json:"pct"`
			} `json:"leagueRecord"`
			Score int `json:"score"`
			Team  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			IsWinner     bool `json:"isWinner"`
			SplitSquad   bool `json:"splitSquad"`
			SeriesNumber int  `json:"seriesNumber"`
		} `json:"home"`
	} `json:"teams"`
	Venue struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue"`
	Content struct {
		Link string `json:"link"`
	} `json:"content"`
	IsTie                  bool   `json:"isTie"`
	GameNumber             int    `json:"gameNumber"`
	PublicFacing           bool   `json:"publicFacing"`
	DoubleHeader           string `json:"doubleHeader"`
	GamedayType            string `json:"gamedayType"`
	Tiebreaker             string `json:"tiebreaker"`
	CalendarEventID        string `json:"calendarEventID"`
	SeasonDisplay          string `json:"seasonDisplay"`
	DayNight               string `json:"dayNight"`
	ScheduledInnings       int    `json:"scheduledInnings"`
	ReverseHomeAwayStatus  bool   `json:"reverseHomeAwayStatus"`
	InningBreakLength      int    `json:"inningBreakLength"`
	GamesInSeries          int    `json:"gamesInSeries"`
	SeriesGameNumber       int    `json:"seriesGameNumber"`
	SeriesDescription      string `json:"seriesDescription"`
	RecordSource           string `json:"recordSource"`
	IfNecessary            string `json:"ifNecessary"`
	IfNecessaryDescription string `json:"ifNecessaryDescription"`
}
