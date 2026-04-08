package database

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/MKW-Limitless-team/limitless-types/ltrc"
	"github.com/MKW-Limitless-team/limitless-types/wwfc"
	_ "github.com/lib/pq"
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
	query := `SELECT player_data.discord_id, player_data.profile_id, mmr, 
				users.profile_id, last_ingamesn, mariokartwii_friend_info, has_ban FROM player_data
				INNER JOIN users on users.profile_id = player_data.profile_id
				AND player_data.discord_id = '%s'`

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
	} else {
		return nil, nil, errors.New("No player found")
	}

	return player, user, nil
}

func GetProfileIDs() []uint64 {
	query := `SELECT profile_id FROM users`

	userIDs := make([]uint64, 0)

	rows, err := globals.GetConnection().Query(query)
	if err != nil {
		return userIDs
	}
	defer rows.Close()

	if rows.Next() {
		var id string
		rows.Scan(&id)
		pid, _ := strconv.Atoi(id)
		userIDs = append(userIDs, uint64(pid))
	}

	return userIDs
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

func RegisterPlayer(discordID string, profileID uint64) error {
	query := `INSERT INTO player_data (discord_id, profile_id)
				VALUES ($1, $2)`

	_, err := globals.GetConnection().Exec(query, discordID, profileID)

	if err != nil {
		return err
	}

	return nil
}
