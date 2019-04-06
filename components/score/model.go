package score

import (
	player "github.com/AndreBtt/Computsal/components/player"
)

type PlayerScore struct {
	Player player.PlayerTable `json:"player"`
	Score  int                `json:"score"`
	Yellow int                `json:"yellowCard"`
	Red    int                `json:"redCard"`
}

type PlayerScoreTable struct {
	ID       int `json:"id"`
	PlayerID int `json:"playerID"`
	MatchID  int `json:"matchID"`
	Quantity int `json:"score"`
	Yellow   int `json:"yellowCard"`
	Red      int `json:"redCard"`
}

type PlayerIDScore struct {
	ID     int `json:"player_id"`
	Score  int `json:"score"`
	Yellow int `json:"yellowCard"`
	Red    int `json:"redCard"`
}
