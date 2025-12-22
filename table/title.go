package table

import (
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/table"
)

const (
	title = "#title"
	Title = "#Title"
)

func titleHandler(table *table.Table, line string) bool {
	title := strings.Replace(line, "#title", "", -1)
	table.Title = removeLeadingSpaces(title)

	return true
}
