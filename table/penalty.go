package table

import (
	"strconv"

	"github.com/MKW-Limitless-team/limitless-types/table"
)

const (
	penalty = "penalty"
	Penalty = "Penalty"
)

func penaltyHandler(table *table.Table, line string) bool {
	penaltyMatch := negativeScoreRegex.FindAllString(line, -1)

	if len(penaltyMatch) != 0 {
		penalty, err := strconv.Atoi(penaltyMatch[0])
		if err == nil {
			table.SetPenalty(penalty)
		}
	}

	return true
}
