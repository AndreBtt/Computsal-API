package nextmatch

type NextMatch struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Type  int    `json:"type"`
}

type NextMatchUpdate struct {
	ID    int    `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Type  int    `json:"type"`
	Time  int    `json:"time"`
}

type NextMatchList struct {
	ID    int    `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Type  int    `json:"type"`
	Time  string `json:"time"`
}

type NextMatchTable struct {
	ID    int    `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Type  int    `json:"type"`
	Time  int    `json:"time"`
}

type NextMatchCreate struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
}

type EliminationMatchTable struct {
	ID    int    `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Type  int    `json:"type"`
}
