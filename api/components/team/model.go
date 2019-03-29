package team

import (
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

type Team struct {
	Name         string
	PhotoURL     string
	Group        int
	Position     int
	Win          int
	Lose         int
	Draw         int
	YellowCard   int
	RedCard      int
	GoalsPro     int
	GoalsAgainst int
	NextMatch    string
	Players      []player.PlayerTable
}
