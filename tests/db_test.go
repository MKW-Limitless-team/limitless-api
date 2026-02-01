package tests

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/nwoik/Limitless-API/db/wwfc"
	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nwoik"
	dbname   = "wwfc"
)

func TestPSQL(t *testing.T) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	t.Run("select users", func(t *testing.T) {
		query := `SELECT profile_id, last_ingamesn, mariokartwii_friend_info FROM users`
		rows, err := db.Query(query)
		if err != nil {
			t.Fatal(err)
		}
		defer rows.Close()

		users := make([]*wwfc.User, 0)

		for rows.Next() {
			user := &wwfc.User{}

			rows.Scan(&user.ProfileId, &user.LastInGameSn, &user.FriendInfo)
			users = append(users, user)
			fmt.Printf("Name: %s\n", user.LastInGameSn)
			fmt.Printf("FC: %d\n", wwfc.PidToFC(user.ProfileId))
			fmt.Printf("PID: %d\n", wwfc.FCToPid(wwfc.PidToFC(user.ProfileId)))
			fmt.Println("--------------------------")
		}

		assert.NotEqual(t, 0, len(users))
	})

	t.Run("", func(t *testing.T) {

	})

}
