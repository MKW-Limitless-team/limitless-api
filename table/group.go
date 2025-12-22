package table

type Group struct {
	Name    string    `json:"name,omitempty"`
	Desc    string    `json:"desc,omitempty"`
	Players []*Player `json:"players,omitempty"`
	Penalty int       `json:"penalty,omitempty"`
	Color   string    `json:"color,omitempty"`
}

func (group *Group) AddPlayer(player *Player) {
	group.Players = append(group.Players, player)
}
