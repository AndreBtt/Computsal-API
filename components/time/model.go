package time

type TimeCreate struct {
	Time string `json: "time"`
}

type TimeDelete struct {
	ID int `json: "id"`
}

type TimeTable struct {
	ID   int    `json:"id"`
	Time string `json:"time"`
}

// Action 0 if want to delete time
// Action 1 if want to update time
type TimeUpdate struct {
	ID     int    `json:"id"`
	Time   string `json: "time"`
	Action int    `json: "action"`
}
