package team

type Team struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo"`
	Group    int    `json:"group_number"`
}
