package main

import (
	"log"
	"net/http"

	"github.com/nwoik/Limitless-API/handlers"
	"github.com/nwoik/Limitless-API/table"
)

func main() {
	table.LoadKeywords()
	http.HandleFunc("/table", handlers.TableHandler)

	log.Print("Running server on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
