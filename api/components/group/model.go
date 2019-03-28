package group

type GroupList struct {
	Number int `json:"group_number"`
}

// Action 0 if want to take the team off the group
// Action 1 if wnat to add the team on the group
type GroupUpdateTeam struct {
	Name   string `json:"team_name"`
	Action int    `json:"action"`
}

type GroupCreate struct {
	Name string `json:"team_name"`
}
