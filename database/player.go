package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nwoik/Limitless-API/database/ltrc"
	"github.com/nwoik/Limitless-API/database/wwfc"
	"github.com/nwoik/Limitless-API/globals"
)

func GetPlayerData(discordID string) (*ltrc.PlayerData, error) {
	query := `SELECT profile_id, discord_id, mmr FROM player_data WHERE discord_id = '%s'`

	rows, err := globals.GetConnection().Query(fmt.Sprintf(query, discordID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	player := &ltrc.PlayerData{}

	if rows.Next() {
		rows.Scan(&player.ProfileID, &player.DiscordID, &player.Mmr)
	}

	return player, nil
}

func GetPlayerInfo(discordID string) (*ltrc.PlayerData, *wwfc.User, error) {
	query := `select discord_id, player_data.profile_id, mmr, 
				users.profile_id, last_ingamesn, mariokartwii_friend_info, has_ban from player_data
				inner join users on users.profile_id = player_data.profile_id
				and discord_id = '%s'`

	rows, err := globals.GetConnection().Query(fmt.Sprintf(query, discordID))
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	player := &ltrc.PlayerData{}
	user := &wwfc.User{}

	if rows.Next() {
		rows.Scan(&player.DiscordID, &player.ProfileID, &player.Mmr,
			&user.ProfileID, &user.LastInGameSn, &user.FriendInfo, &user.HasBan)
	}

	return player, user, nil
}

func GetUser(profileID uint64) (*wwfc.User, error) {
	query := `SELECT profile_id, last_ingamesn, mariokartwii_friend_info, has_ban FROM users WHERE profile_id = %d`
	rows, err := globals.GetConnection().Query(fmt.Sprintf(query, profileID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &wwfc.User{}

	if rows.Next() {
		rows.Scan(&user.ProfileID, &user.LastInGameSn, &user.FriendInfo, &user.HasBan)
	}

	return user, nil
}
