package table

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	trailingLinesRegex, _ = regexp.Compile(`\s+$`)
	flagRegex, _          = regexp.Compile(`\[\w+\]`)
	scoresRegex, _        = regexp.Compile(`\b[\d+\-|]+`)
	negativeScoreRegex, _ = regexp.Compile(`-\d+`)
)

func TokenizePlayer(sample string) *Player {
	player := &Player{}
	sample, penalty := getPenalty(sample)
	player.Penalty = penalty

	flagMatch := flagRegex.FindAllString(sample, -1)
	sample = flagRegex.ReplaceAllString(sample, "")

	scoresMatch := scoresRegex.FindAllString(sample, -1)
	sample = scoresRegex.ReplaceAllString(sample, "")

	player.Name = trailingLinesRegex.ReplaceAllLiteralString(sample, "")

	if len(flagMatch) != 0 {
		player.Flag = flagMatch[0]
	}

	if len(scoresMatch) != 0 {
		scores := strings.Split(scoresMatch[0], "|")

		for _, score := range scores {
			addScore(player, score)
		}
	}

	return player
}

func addScore(player *Player, score string) {
	if strings.Contains(score, "+") {
		player.Scores = append(player.Scores, sumScores(strings.Split(score, "+")))
	} else {
		scoreNumber, err := strconv.Atoi(score)
		if err == nil {
			player.Scores = append(player.Scores, scoreNumber)
		}
	}
}

func getPenalty(sample string) (string, int) {
	penaltyMatch := negativeScoreRegex.FindAllString(sample, -1)
	penalty := 0

	for _, penaltyString := range penaltyMatch {
		penaltyValue, err := strconv.Atoi(penaltyString)
		if err == nil {
			penalty += penaltyValue
			sample = strings.Replace(sample, penaltyString, "", 1)
		}
	}

	return sample, penalty
}

func sumScores(sample []string) int {
	result := 0

	for _, numString := range sample {
		num, err := strconv.Atoi(numString)
		if err == nil {
			result += num
		}
	}

	return result
}
