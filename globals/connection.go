package globals

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB       *sql.DB
	password = os.Getenv("DB_PASS")
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "wwfc"
)

func Initialize() error {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Print("Successfully connected to DB!")

	DB = db

	return nil
}

func GetConnection() *sql.DB {
	err := DB.Ping()
	if err != nil {
		return DB
	}

	Initialize()
	return DB
}
