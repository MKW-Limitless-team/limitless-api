package table

type Player struct {
	Name    string `json:"name,omitempty"`
	Flag    string `json:"flag,omitempty"`
	Scores  []int  `json:"scores,omitempty"`
	Penalty int    `json:"penalty,omitempty"`
}
