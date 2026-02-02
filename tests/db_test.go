package tests

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/nwoik/Limitless-API/database/ltrc"
	"github.com/nwoik/Limitless-API/database/wwfc"
	"github.com/nwoik/Limitless-API/globals"
	"github.com/stretchr/testify/assert"
)

func TestPSQL(t *testing.T) {
	globals.Initialize()

	t.Run("get user", func(t *testing.T) {
		query := `SELECT profile_id, last_ingamesn, mariokartwii_friend_info FROM users`
		rows, err := globals.GetConnection().Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		users := make([]*wwfc.User, 0)

		for rows.Next() {
			user := &wwfc.User{}

			rows.Scan(&user.ProfileID, &user.LastInGameSn, &user.FriendInfo)
			users = append(users, user)
		}

		assert.NotEqual(t, 0, len(users))
	})

	t.Run("get playerdata", func(t *testing.T) {
		query := `SELECT profile_id, discord_id, mmr FROM player_data`

		rows, err := globals.GetConnection().Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		players := make([]*ltrc.PlayerData, 0)

		for rows.Next() {
			player := &ltrc.PlayerData{}

			rows.Scan(&player.ProfileID, &player.DiscordID, &player.Mmr)
			players = append(players, player)
		}

		assert.NotEqual(t, 0, len(players))
	})

}
