package main

import (
	"log"
	"net/http"

	"github.com/MKW-Limitless-team/limitless-api/db"
	"github.com/MKW-Limitless-team/limitless-api/handlers"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	_ "github.com/ncruces/go-sqlite3/vfs/memdb"
)

func main() {
	err := db.Initialize("./ltrc.db")
	http.HandleFunc("/api/player", handlers.GetPlayerHandler)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print("Running server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
