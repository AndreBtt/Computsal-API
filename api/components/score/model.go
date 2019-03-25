package score

import (
	player "github.com/AndreBtt/Computsal/api/components/player"
)

type PlayerScore struct {
	Player player.PlayerTable
	Score  int
}

type PlayerScoreTable struct {
	ID       int
	PlayerID int
	MatchID  int
	Quantity int
}
