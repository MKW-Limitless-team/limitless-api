package tests

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/nwoik/Limitless-API/database"
	"github.com/nwoik/Limitless-API/database/ltrc"
	"github.com/nwoik/Limitless-API/database/wwfc"
	"github.com/nwoik/Limitless-API/globals"
	"github.com/stretchr/testify/assert"
)

func TestPSQL(t *testing.T) {
	globals.Initialize()

	t.Run("select", func(t *testing.T) {
		t.Run("get users", func(t *testing.T) {
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

		t.Run("get players data", func(t *testing.T) {
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
	})

	t.Run("player", func(t *testing.T) {
		t.Run("get player data", func(t *testing.T) {
			player, err := database.GetPlayerData("593475382498426927")

			if err != nil {
				t.Fatal(err)
			}

			assert.NotEqual(t, nil, player)
			assert.Equal(t, uint64(436787076), player.ProfileID)
		})

		t.Run("get user", func(t *testing.T) {
			user, err := database.GetUser(436787076)

			if err != nil {
				t.Fatal(err)
			}

			assert.NotEqual(t, nil, user)
			assert.NotEqual(t, "", user.LastInGameSn)
		})

		t.Run("get player info", func(t *testing.T) {
			player, user, err := database.GetPlayerInfo("593475382498426927")

			if err != nil {
				t.Fatal(err)
			}

			assert.NotEqual(t, nil, player)
			assert.Equal(t, uint64(436787076), player.ProfileID)
			assert.NotEqual(t, nil, user)
			assert.NotEqual(t, "", user.LastInGameSn)
		})
	})
}
