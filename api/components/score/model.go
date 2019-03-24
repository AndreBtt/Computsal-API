package score

import (
	player "github.com/AndreBtt/Computsal/api/components/player"
)

type PlayerScore struct {
	Player player.PlayerTable
	Score  int
}
