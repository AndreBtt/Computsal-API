package team

import (
	"github.com/AndreBtt/Computsal/api/components/match"
	"github.com/AndreBtt/Computsal/api/components/player"
)

type TeamTable struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo"`
	Group    int    `json:"group"`
}

type TeamUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo"`
}

type TeamCreate struct {
	Name         string              `json:"name"`
	PhotoURL     string              `json:"photo"`
	Players      []player.PlayerName `json:"players"`
	CaptainEmail string              `json:"captain_email"`
}

type TeamNextMatch struct {
	Name string `json:"name"`
	Time int    `json:"time"`
}

type Team struct {
	ID              int
	Name            string
	PhotoURL        string
	Group           int
	Win             int
	Lose            int
	Draw            int
	GoalsPro        int
	GoalsAgainst    int
	NextMatch       TeamNextMatch
	CaptainName     string
	Players         []player.PlayerTable
	PreviousMatches []match.PreviousMatchList
}
