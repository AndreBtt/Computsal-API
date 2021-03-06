package group

import "github.com/AndreBtt/Computsal/components/team"

type GroupList struct {
	Number int `json:"group_number"`
}

// Action 0 if want to take the team off the group
// Action 1 if want to add the team on the group
type GroupUpdateTeam struct {
	ID     int `json:"id"`
	Action int `json:"action"`
}

type Group struct {
	Number int              `json:"group_number"`
	Team   []team.TeamGroup `json:"teams"`
}
