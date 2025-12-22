package table

import (
	"strings"
)

var (
	Keywords = map[string]func(table *Table, line string) bool{}
)

func LoadKeywords() {
	Keywords[penalty] = penaltyHandler
	Keywords[Penalty] = penaltyHandler
	Keywords[title] = titleHandler
	Keywords[Title] = titleHandler
}

func handleKeyword(table *Table, line string) bool {
	for key, valueFunc := range Keywords {
		if strings.HasPrefix(line, key) {
			return valueFunc(table, line)
		}
	}

	return false
}
