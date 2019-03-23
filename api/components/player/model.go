package player

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Team string `json:"fk_player_team"`
}
