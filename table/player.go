package table

import "regexp"

var (
	playerNameRegex, _    = regexp.Compile(`^\w+`)
	flagRegex, _          = regexp.Compile(`\[\w+\]`)
	scoresRegex, _        = regexp.Compile(`[\d+-|]+$`)
	negativeScoreRegex, _ = regexp.Compile(`-\d+`)
)

type Player struct {
	Name    string
	Flag    string
	Scores  []int
	Penalty int
}
