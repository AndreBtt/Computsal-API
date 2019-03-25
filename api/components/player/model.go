package player

type PlayerTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
}

type Player struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Team         string `json:"team"`
	TeamPhotoURL string `json:"teamPhoto"`
	Score        int    `json:"score"`
	YellowCard   int    `json:"yellowCard"`
	RedCard      int    `json:"redCard"`
	Captain      bool   `json:"captain"`
}
