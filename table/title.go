package table

import (
	"strings"
)

const (
	title = "#title"
	Title = "#Title"
)

func titleHandler(table *Table, line string) bool {
	title := strings.Replace(line, "#title", "", -1)
	table.Title = removeLeadingSpaces(title)

	return true
}
