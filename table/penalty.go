package table

import (
	"strconv"
)

const (
	penalty = "penalty"
	Penalty = "Penalty"
)

func penaltyHandler(table *Table, line string) bool {
	penaltyMatch := negativeScoreRegex.FindAllString(line, -1)

	if len(penaltyMatch) != 0 {
		penalty, err := strconv.Atoi(penaltyMatch[0])
		if err == nil {
			table.SetPenalty(penalty)
		}
	}

	return true
}
