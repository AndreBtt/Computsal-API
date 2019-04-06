package schedule

type TimeAvailable struct {
	TimeID       int    `json:"time_id"`
	Time         string `json:"time"`
	Availability bool   `json:"availability`
}

type TimeUpdate struct {
	TimeID       int  `json:"time_id"`
	Availability bool `json:"availability`
}
