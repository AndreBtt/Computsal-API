package team

import (
	player "github.com/AndreBtt/Computsal/api/components/player"
)

type TeamTable struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo"`
	Group    int    `json:"group_number"`
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
	NextGame     string
	Players      []player.PlayerTable
}
