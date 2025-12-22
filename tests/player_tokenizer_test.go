package tests

import (
	"testing"

	"github.com/nwoik/Limitless-API/table"
	"github.com/stretchr/testify/assert"
)

func TestPlayerTokenizer(t *testing.T) {
	t.Run("player name and flag", func(t *testing.T) {
		playerString := "Billy [gb]"

		player := table.TokenizePlayer(playerString)

		assert.Equal(t, player.Name, "Billy")
		assert.Equal(t, player.Flag, "[gb]")
	})

	t.Run("player basic score", func(t *testing.T) {
		playerString := "Billy [gb] 110"

		player := table.TokenizePlayer(playerString)
		scores := []int{110}

		assert.Equal(t, player.Scores, scores)
	})

	t.Run("player basic score and penalty", func(t *testing.T) {
		playerString := "Billy [gb] 110-10"

		player := table.TokenizePlayer(playerString)
		scores := []int{110}
		penalty := -10

		assert.Equal(t, player.Scores, scores)
		assert.Equal(t, player.Penalty, penalty)
	})

	t.Run("player score addition", func(t *testing.T) {
		playerString := "Billy [gb] 110+10"

		player := table.TokenizePlayer(playerString)
		scores := []int{120}

		assert.Equal(t, player.Scores, scores)
	})

	t.Run("player multiple score", func(t *testing.T) {
		playerString := "Billy [gb] 110|67|40"

		player := table.TokenizePlayer(playerString)
		scores := []int{110, 67, 40}

		assert.Equal(t, player.Scores, scores)
	})

	t.Run("player multiple score with addition", func(t *testing.T) {
		playerString := "Billy [gb] 110|67+3|40"

		player := table.TokenizePlayer(playerString)
		scores := []int{110, 70, 40}

		assert.Equal(t, player.Scores, scores)
	})

	t.Run("player multiple score with addition and penalty", func(t *testing.T) {
		playerString := "Billy [gb] -10+110-10+4|67+3|40-10"

		player := table.TokenizePlayer(playerString)
		scores := []int{114, 70, 40}
		penalty := -30

		assert.Equal(t, player.Scores, scores)
		assert.Equal(t, player.Penalty, penalty)
	})
}
