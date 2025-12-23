package tests

import (
	"strings"
	"testing"

	"github.com/nwoik/Limitless-API/table"
	"github.com/stretchr/testify/assert"
)

func TestTableTokenizer(t *testing.T) {
	table.LoadKeywords()
	t.Run("table title", func(t *testing.T) {
		sample := `
			#title 4 races
		`

		table := table.ProcessTable(stripTabs(sample))

		assert.Equal(t, table.Title, "4 races")
	})

	t.Run("table group name", func(t *testing.T) {
		sample := `
			FFA
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Name, "FFA")
		assert.Equal(t, table.GetCurrentGroup().Desc, "")
	})

	t.Run("table group desc", func(t *testing.T) {
		sample := `
			-free for all
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Name, "")
		assert.Equal(t, table.GetCurrentGroup().Desc, "free for all")
	})

	t.Run("table group color", func(t *testing.T) {
		sample := `
			#ff0000
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Color, "#ff0000")
	})

	t.Run("table group penalty", func(t *testing.T) {
		sample := `
			penalty -10
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Penalty, -10)
	})

	t.Run("table group Penalty", func(t *testing.T) {
		sample := `
			Penalty -10
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Penalty, -10)
	})

	t.Run("table group name and empty desc", func(t *testing.T) {
		sample := `
			FFA-
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Name, "FFA")
		assert.Equal(t, table.GetCurrentGroup().Desc, "")
	})

	t.Run("table group name, desc and color", func(t *testing.T) {
		sample := `
			FFA-free for all #ff0000
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.GetCurrentGroup().Name, "FFA")
		assert.Equal(t, table.GetCurrentGroup().Color, "#ff0000")
		assert.Equal(t, table.GetCurrentGroup().Desc, "free for all")
	})

	t.Run("FFA table", func(t *testing.T) {
		sample := `
			#title 4 races
			FFA-free for all #ff0000
			Alice [us] 112
			Billy [gb] 110
			Carol [au] 76
			Derek [ca] 72
			Ellen [de] 90-10
			Frank [ie] 55
			Grant [cl] 70+20+8
			Henry [br] 78
			Isaac 46
			James [kr] 100
			Karen [jp] 68
			Lucas [mx] 79
		`

		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.Title, "4 races")
		assert.Equal(t, table.GetCurrentGroup().Name, "FFA")
		assert.Equal(t, table.GetCurrentGroup().Desc, "free for all")
		assert.Equal(t, table.GetCurrentGroup().Color, "#ff0000")
		assert.Equal(t, len(table.Groups[0].Players), 12)
	})

	t.Run("Team table", func(t *testing.T) {
		sample := `
			A - Full Team Name #ff0000
			Alice [us] 112|65|42
			Billy [gb] 110|32|88
			Carol [au] 76|18|45
			Derek [ca] 72|26|79
			Ellen [de] 90-10|80|54
			Frank [ie] 55|38|34
			Penalty -10

			B
			Grant [cl] 70+20+8|50|62
			Henry [br] 78|45|70
			Isaac 46|28|61
			James [kr] 100|80|49
			Karen [jp] 68|36|38
			Lucas [mx] 79|15|108
		`
		table := table.ProcessTable(stripTabs(sample))
		assert.Equal(t, table.Groups[0].Name, "A")
		assert.Equal(t, table.Groups[0].Desc, "Full Team Name")
		assert.Equal(t, table.Groups[0].Color, "#ff0000")
		assert.Equal(t, table.Groups[0].Penalty, -10)
		assert.Equal(t, len(table.Groups[0].Players), 6)

		assert.Equal(t, table.Groups[1].Name, "B")
		assert.Equal(t, len(table.Groups[1].Players), 6)

	})
}

func stripTabs(sample string) string {
	return strings.ReplaceAll(sample, "\t", "")
}
