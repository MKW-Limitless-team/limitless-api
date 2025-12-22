package main

import (
	"log"
	"net/http"

	"github.com/nwoik/Limitless-API/handlers"
)

func main() {
	http.HandleFunc("/table", handlers.TableHandler)

	log.Print("Running server on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
