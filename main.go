package main

import (
	"log"
	"net/http"

	"github.com/nwoik/Limitless-API/globals"
	"github.com/nwoik/Limitless-API/handlers"
	"github.com/nwoik/Limitless-API/table"
)

func main() {
	globals.Initialize()
	defer globals.GetConnection().Close()

	table.LoadKeywords()
	http.HandleFunc("/table", handlers.TableHandler)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/player", handlers.Player)

	log.Print("Running server on http://localhost:5000")
	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
