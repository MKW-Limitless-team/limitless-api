package ltrc

import (
	"sort"
)

type PlayerData struct {
	ProfileID uint64 `json:"profileID,omitempty"`
	DiscordID string `json:"discordID,omitempty"`
	Mmr       *int64 `json:"mmr,omitempty"`
}

type byMmr []*PlayerData

func (b byMmr) Len() int {
	return len(b)
}

func (b byMmr) Less(i int, j int) bool {
	return *b[i].Mmr > *b[j].Mmr
}

func (b byMmr) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func SortByMMR(players []*PlayerData) []*PlayerData {
	sort.Sort(byMmr(players))

	return players
}
