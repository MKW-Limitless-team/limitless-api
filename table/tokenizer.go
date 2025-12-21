package table

import (
	"strconv"
	"strings"
)

func TokenizePlayer(sample string) *Player {
	player := &Player{}
	sample, penalty := getPenalty(sample)
	player.Penalty = penalty

	playerNameMatch := playerNameRegex.FindAllString(sample, -1)
	flagMatch := flagRegex.FindAllString(sample, -1)
	scoresMatch := scoresRegex.FindAllString(sample, -1)

	if len(playerNameMatch) != 0 {
		player.Name = playerNameMatch[0]
	}

	if len(flagMatch) != 0 {
		player.Flag = flagMatch[0]
	}

	if len(scoresMatch) != 0 {
		scores := strings.Split(scoresMatch[0], "|")

		for _, score := range scores {
			if strings.Contains(score, "+") {
				player.Scores = append(player.Scores, addScores(strings.Split(score, "+")))
			} else {
				scoreNumber, _ := strconv.Atoi(score)
				player.Scores = append(player.Scores, scoreNumber)
			}
		}
	}

	return player
}

func getPenalty(sample string) (string, int) {
	penaltyMatch := negativeScoreRegex.FindAllString(sample, -1)
	penalty := 0

	for _, penaltyString := range penaltyMatch {
		penaltyValue, _ := strconv.Atoi(penaltyString)
		penalty += penaltyValue
		sample = strings.Replace(sample, penaltyString, "", 1)
	}

	return sample, penalty
}

func addScores(sample []string) int {
	result := 0

	for _, numString := range sample {
		num, _ := strconv.Atoi(numString)
		result += num
	}

	return result
}
