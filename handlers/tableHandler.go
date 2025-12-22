package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	t "github.com/nwoik/Limitless-API/table"
)

func TableHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	log.Println("Handling table request")

	value := r.URL.Query().Get("data")
	table := t.ProcessTable(value)

	resp, err := json.Marshal(table)
	if err != nil {
		log.Println(err)
	}

	w.Write(resp)
	log.Println("Sent table response")
}
