package tests

import (
	"strings"
	"testing"

	"github.com/nwoik/Limitless-API/table"
	"github.com/stretchr/testify/assert"
)

func TestTableTokenizer(t *testing.T) {
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
}

func stripTabs(sample string) string {
	return strings.ReplaceAll(sample, "\t", "")
}
