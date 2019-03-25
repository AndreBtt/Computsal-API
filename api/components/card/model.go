package card

import "github.com/AndreBtt/Computsal/api/components/player"

type PlayerCard struct {
	Player player.PlayerTable
	Red    int
	Yellow int
}

type CardTable struct {
	ID       int
	PlayerID int
	MatchID  int
	Yellow   int
	Red      int
}

type CardUpdate struct {
	ID     int
	Yellow int
	Red    int
}
