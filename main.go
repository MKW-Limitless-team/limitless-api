package main

import (
	"log"
	"net/http"
)

func main() {

	log.Print("Running server on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
