package team

import (
	"github.com/AndreBtt/Computsal/components/player"
	"github.com/AndreBtt/Computsal/components/previousmatch"
)

type TeamTable struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo"`
	Group    int    `json:"group"`
}

type TeamUpdate struct {
	ID       int                   `json:"id"`
	Name     string                `json:"name"`
	PhotoURL string                `json:"photo"`
	Players  []player.PlayerUpdate `json:"players"`
}

type TeamCreate struct {
	Name         string              `json:"name"`
	PhotoURL     string              `json:"photo"`
	Players      []player.PlayerName `json:"players"`
	CaptainEmail string              `json:"captain_email"`
}

type TeamNextMatch struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

type Team struct {
	ID              int                               `json:"id"`
	Name            string                            `json:"name"`
	PhotoURL        string                            `json:"photo"`
	Group           int                               `json:"group"`
	Win             int                               `json:"win"`
	Lose            int                               `json:"lose"`
	Draw            int                               `json:"draw"`
	GoalsPro        int                               `json:"goals_pro"`
	GoalsAgainst    int                               `json:"goals_against"`
	NextMatch       TeamNextMatch                     `json:"next_match"`
	CaptainName     string                            `json:"captain"`
	Players         []player.PlayerTeamScore          `json:"players"`
	PreviousMatches []previousmatch.PreviousMatchList `json:"previous_matches"`
}

type TeamGroup struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PhotoURL     string `json:"photo"`
	Win          int    `json:"win"`
	Lose         int    `json:"lose"`
	Draw         int    `json:"draw"`
	GoalsPro     int    `json:"goals_pro"`
	GoalsAgainst int    `json:"goals_against"`
	Points       int    `json:"points"`
}
