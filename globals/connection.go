package globals

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nwoik"
	dbname   = "wwfc"
)

func Initialize() error {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
