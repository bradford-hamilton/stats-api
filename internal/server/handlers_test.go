package server

import (
	"reflect"
	"testing"
	"time"
)

func Test_sortGames(t *testing.T) {
	type args struct {
		games  []Game
		teamID int
	}
	tests := []struct {
		name        string
		description string
		args        args
		want        []Game
	}{
		{
			name:        "Sort - no favorite team match",
			description: "When there is no favorite teamID match, the original order is preserved and returned",
			args: args{
				games:  fixture1,
				teamID: 11,
			},
			want: fixture1,
		},
		{
			name:        "Sort - favorite team match",
			description: "When a favorite teamID is matched, that team is placed in the front of the list while preserving order otherwise",
			args: args{
				games:  fixture1,
				teamID: 143,
			},
			want: fixture1SortedByPhilliesID,
		},
		{
			name:        "Sort - favorite team match double header 'Y'",
			description: "When a favorite teamID is matched and that team has a double header type 'Y' (single admission/traditional), the two games are moved to the front of the list. If the first favorite team game has a 'startTimeTBD' it is moved into the 2nd game position",
			args: args{
				games:  fixture2,
				teamID: 143,
			},
			want: fixture2SortedByPhilliesID,
		},
		{
			name:        "Sort - favorite team match double header 'S'",
			description: "When a favorite teamID is matched and that team has a double header type 'S' (split admission), the two games are moved to the front of the list. If the first favorite team game has a a later 'gameDate' it is moved into the 2nd game position",
			args: args{
				games:  fixture3,
				teamID: 143,
			},
			want: fixture3SortedByPhilliesID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortGames(tt.args.games, tt.args.teamID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

var fixture1 = []Game{{
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 115, Name: "Colorado Rockies"}},
		Home: Team{Team: TeamMetadata{ID: 134, Name: "Pittsburgh Pirates"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 144, Name: "Atlanta Braves"}},
		Home: Team{Team: TeamMetadata{ID: 147, Name: "New York Yankees"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "N",
}}

var fixture1SortedByPhilliesID = []Game{{
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 115, Name: "Colorado Rockies"}},
		Home: Team{Team: TeamMetadata{ID: 134, Name: "Pittsburgh Pirates"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 144, Name: "Atlanta Braves"}},
		Home: Team{Team: TeamMetadata{ID: 147, Name: "New York Yankees"}},
	},
	DoubleHeader: "N",
}}

var fixture2 = []Game{{
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 115, Name: "Colorado Rockies"}},
		Home: Team{Team: TeamMetadata{ID: 134, Name: "Pittsburgh Pirates"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: true},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "Y",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 144, Name: "Atlanta Braves"}},
		Home: Team{Team: TeamMetadata{ID: 147, Name: "New York Yankees"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "Y",
}}

var fixture2SortedByPhilliesID = []Game{{
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "Y",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: true},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "Y",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 115, Name: "Colorado Rockies"}},
		Home: Team{Team: TeamMetadata{ID: 134, Name: "Pittsburgh Pirates"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 144, Name: "Atlanta Braves"}},
		Home: Team{Team: TeamMetadata{ID: 147, Name: "New York Yankees"}},
	},
	DoubleHeader: "N",
}}

var fixture3gameDateEarlier = time.Now()
var fixture3gameDateLater = time.Now().Add(2 * time.Hour)

var fixture3 = []Game{{
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 115, Name: "Colorado Rockies"}},
		Home: Team{Team: TeamMetadata{ID: 134, Name: "Pittsburgh Pirates"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: fixture3gameDateLater,
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "S",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 144, Name: "Atlanta Braves"}},
		Home: Team{Team: TeamMetadata{ID: 147, Name: "New York Yankees"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: fixture3gameDateEarlier,
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "S",
}}

var fixture3SortedByPhilliesID = []Game{{
	GameDate: fixture3gameDateEarlier,
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "S",
}, {
	GameDate: fixture3gameDateLater,
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 143, Name: "Philadelphia Phillies"}},
		Home: Team{Team: TeamMetadata{ID: 146, Name: "Miami Marlins"}},
	},
	DoubleHeader: "S",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 115, Name: "Colorado Rockies"}},
		Home: Team{Team: TeamMetadata{ID: 134, Name: "Pittsburgh Pirates"}},
	},
	DoubleHeader: "N",
}, {
	GameDate: time.Time{},
	Status:   Status{DetailedState: "Final", StartTimeTBD: false},
	Teams: Teams{
		Away: Team{Team: TeamMetadata{ID: 144, Name: "Atlanta Braves"}},
		Home: Team{Team: TeamMetadata{ID: 147, Name: "New York Yankees"}},
	},
	DoubleHeader: "N",
}}
