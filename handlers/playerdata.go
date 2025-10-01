package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MKW-Limitless-team/limitless-api/db"
)

func GetPlayerHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	playerData, err := db.GetPlayer("593475382498426927")

	if err != nil {
		log.Print(err)
	}

	response, err := json.Marshal(&playerData)
	if err != nil {
		log.Print(err)
	}

	log.Print(string(response))
	w.Write(response)
}
