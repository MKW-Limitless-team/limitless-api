package table

import (
	"fmt"
	"time"
)

type Table struct {
	Title  string   `json:"title,omitempty"`
	Groups []*Group `json:"groups,omitempty"`
	Date   string   `json:"date,omitempty"`
}

func NewTable() *Table {
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("%d %s %d", day, month, year)
	table := &Table{
		Date: date,
	}

	return table
}

func (table *Table) AddGroup(group *Group) {
	table.Groups = append(table.Groups, group)
}

func (table *Table) SetPenalty(penalty int) {
	table.GetCurrentGroup().Penalty = penalty
}

func (table *Table) AddPlayer(player *Player) {
	table.GetCurrentGroup().AddPlayer(player)
}

func (table *Table) GetCurrentGroup() *Group {
	if len(table.Groups) == 0 {
		table.AddGroup(&Group{})
	}
	return table.Groups[len(table.Groups)-1]
}
