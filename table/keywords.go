package table

import (
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/table"
)

var (
	Keywords = map[string]func(table *table.Table, line string) bool{}
)

func LoadKeywords() {
	Keywords[penalty] = penaltyHandler
	Keywords[Penalty] = penaltyHandler
	Keywords[title] = titleHandler
	Keywords[Title] = titleHandler
}

func handleKeyword(table *table.Table, line string) bool {
	for key, valueFunc := range Keywords {
		if strings.HasPrefix(line, key) {
			return valueFunc(table, line)
		}
	}

	return false
}
