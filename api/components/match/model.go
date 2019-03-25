package match

import (
	score "github.com/AndreBtt/Computsal/api/components/score"
)

type PreviousMatchQuery struct {
	ID    int
	Type  int
	Phase int
	Team1 string
	Team2 string
	Team  string
	Score int
}

type PreviousMatchList struct {
	ID     int    `json:"id"`
	Team1  string `json:"team1"`
	Team2  string `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
	Type   int    `json:"type"`
	Phase  int    `json:"phase"`
}

type PreviousMatch struct {
	ID          int                 `json:"id"`
	Team1       string              `json:"team1"`
	Team2       string              `json:"team2"`
	YellowCard1 int                 `json:"yellowCard1"`
	YellowCard2 int                 `json:"yellowCard2"`
	RedCard1    int                 `json:"redCard1"`
	RedCard2    int                 `json:"redCard2"`
	Score1      int                 `json:"score1"`
	Score2      int                 `json:"score2"`
	Type        int                 `json:"type"`
	Group       int                 `json:"group"`
	PlayerScore []score.PlayerScore `json:"player"`
}
