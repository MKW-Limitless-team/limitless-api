package table

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	trailingSpacesRegex, _ = regexp.Compile(`\s+$`)
	leadingSpacesRegex, _  = regexp.Compile(`^\s+`)
	flagRegex, _           = regexp.Compile(`\[\w+\]`)
	scoresRegex, _         = regexp.Compile(`[\d+|\-]+[\d+|\-]+`)
	negativeScoreRegex, _  = regexp.Compile(`-\d+`)
	colorRegex, _          = regexp.Compile(`#[a-fA-F0-9]{6}`)
)

func ProcessTable(sample string) *Table {
	table := NewTable()
	lines := strings.Split(sample, "\n")

	for _, line := range lines {
		if removeTralingSpaces(line) == "" {

		} else if strings.HasPrefix(line, "#title") {
			title := strings.Replace(line, "#title", "", -1)
			table.Title = removeLeadingSpaces(title)
		} else if isPenalty(line) {
			penaltyMatch := negativeScoreRegex.FindAllString(line, -1)

			if len(penaltyMatch) != 0 {
				penalty, err := strconv.Atoi(penaltyMatch[0])
				if err == nil {
					table.SetPenalty(penalty)
				}
			}
		} else if isPlayer(line) && !colorRegex.MatchString(line) {
			player := TokenizePlayer(line)
			table.AddPlayer(player)
		} else {
			group := TokenizeGroup(line)
			table.AddGroup(group)
		}
	}

	return table
}

func TokenizeGroup(sample string) *Group {
	group := &Group{}

	colorMatch := colorRegex.FindAllString(sample, -1)
	sample = removeTralingSpaces(colorRegex.ReplaceAllString(sample, ""))

	if len(colorMatch) != 0 {
		group.Color = colorMatch[0]
	}

	if strings.Contains(sample, "-") {
		titleAndDesc := strings.Split(sample, "-")
		group.Name = titleAndDesc[0]
		group.Desc = titleAndDesc[len(titleAndDesc)-1]
	} else {
		group.Name = sample
	}

	return group
}

func TokenizePlayer(sample string) *Player {
	player := &Player{}
	sample, penalty := getPenalty(sample)
	player.Penalty = penalty

	flagMatch := flagRegex.FindAllString(sample, -1)
	sample = flagRegex.ReplaceAllString(sample, "")

	scoresMatch := scoresRegex.FindAllString(sample, -1)
	sample = scoresRegex.ReplaceAllString(sample, "")

	player.Name = removeTralingSpaces(sample)

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

func removeLeadingSpaces(sample string) string {
	return leadingSpacesRegex.ReplaceAllLiteralString(sample, "")
}

func removeTralingSpaces(sample string) string {
	return trailingSpacesRegex.ReplaceAllLiteralString(sample, "")
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

func isPenalty(sample string) bool {
	return strings.HasPrefix(sample, "Penalty") || strings.HasPrefix(sample, "penalty")
}

func isPlayer(sample string) bool {
	flagMatch := flagRegex.FindAllString(sample, -1)
	scoresMatch := scoresRegex.FindAllString(sample, -1)

	if len(flagMatch) != 0 {
		return true
	}

	if len(scoresMatch) != 0 {
		return true
	}

	return false
}
