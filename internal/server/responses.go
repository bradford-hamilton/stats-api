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
	Dates                []Date `json:"dates"`
}

// Date represents a single date from within a StatScheduleResponse
type Date struct {
	Date                 string        `json:"date"`
	TotalItems           int           `json:"totalItems"`
	TotalEvents          int           `json:"totalEvents"`
	TotalGames           int           `json:"totalGames"`
	TotalGamesInProgress int           `json:"totalGamesInProgress"`
	Games                []Game        `json:"games"`
	Events               []interface{} `json:"events"`
}

// Game represents a single game from within a StatScheduleResponse
type Game struct {
	GamePk                 int       `json:"gamePk"`
	Link                   string    `json:"link"`
	GameType               string    `json:"gameType"`
	Season                 string    `json:"season"`
	GameDate               time.Time `json:"gameDate"`
	OfficialDate           string    `json:"officialDate"`
	Status                 Status    `json:"status"`
	Teams                  Teams     `json:"teams"`
	Venue                  Venue     `json:"venue"`
	Content                Content   `json:"content"`
	IsTie                  bool      `json:"isTie"`
	GameNumber             int       `json:"gameNumber"`
	PublicFacing           bool      `json:"publicFacing"`
	DoubleHeader           string    `json:"doubleHeader"`
	GamedayType            string    `json:"gamedayType"`
	Tiebreaker             string    `json:"tiebreaker"`
	CalendarEventID        string    `json:"calendarEventID"`
	SeasonDisplay          string    `json:"seasonDisplay"`
	DayNight               string    `json:"dayNight"`
	ScheduledInnings       int       `json:"scheduledInnings"`
	ReverseHomeAwayStatus  bool      `json:"reverseHomeAwayStatus"`
	InningBreakLength      int       `json:"inningBreakLength"`
	GamesInSeries          int       `json:"gamesInSeries"`
	SeriesGameNumber       int       `json:"seriesGameNumber"`
	SeriesDescription      string    `json:"seriesDescription"`
	RecordSource           string    `json:"recordSource"`
	IfNecessary            string    `json:"ifNecessary"`
	IfNecessaryDescription string    `json:"ifNecessaryDescription"`
}

// Status represents a single status from within a StatScheduleResponse.
type Status struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
	AbstractGameCode  string `json:"abstractGameCode"`
}

// Teams represents both the away and home teams from within a StatScheduleResponse.
type Teams struct {
	Away Team `json:"away"`
	Home Team `json:"home"`
}

// Team represents team data from within a StatScheduleResponse.
type Team struct {
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int          `json:"score"`
	Team         TeamMetadata `json:"team"`
	IsWinner     bool         `json:"isWinner"`
	SplitSquad   bool         `json:"splitSquad"`
	SeriesNumber int          `json:"seriesNumber"`
}

// LeagueRecord represents a league record from within a StatScheduleResponse.
type LeagueRecord struct {
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Pct    string `json:"pct"`
}

// TeamMetadata represents some more identifying information for a team from within a StatScheduleResponse.
type TeamMetadata struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Venue represents identifying information for venue from within a StatScheduleResponse.
type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Content contains a link to more game content from within a StatScheduleResponse.
type Content struct {
	Link string `json:"link"`
}
