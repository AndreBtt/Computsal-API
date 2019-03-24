package player

type PlayerTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
}

type Player struct {
	ID           int
	Name         string
	Team         string
	TeamPhotoURL string
	Score        int
	YellowCard   int
	RedCard      int
	Captain      bool
}
