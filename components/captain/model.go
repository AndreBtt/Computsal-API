package captain

type CaptainQuery struct {
	PlayerID int    `json:"player_id"`
	Email    string `json:"email"`
}

type CaptainCreate struct {
	Team     string `json:"team"`
	PlayerID int    `json:"player_id"`
	Email    string `json:"email"`
}

type CaptainTeam struct {
	Team string `json:"team"`
}
