package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MKW-Limitless-team/limitless-types/responses"
	"github.com/nwoik/Limitless-API/database"
)

func Player(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var status *responses.PlayerInfoResponse
	discordID := r.URL.Query().Get("discord_id")
	player, user, err := database.GetPlayerInfo(discordID)

	if err != nil {
		status = responses.FailureResponse("Failed to get player")
		resp, _ := json.Marshal(status)
		w.Write(resp)
		log.Println(status)
		return
	}

	status = responses.SuccessResponse()
	status.PlayerData = player
	status.User = user
	resp, _ := json.Marshal(status)
	w.Write(resp)
	log.Println(status)
}

func Register(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var status *responses.PlayerInfoResponse
	discordID := r.URL.Query().Get("discord_id")
	profileID := r.URL.Query().Get("profile_id")

	pid, err := strconv.Atoi(profileID)

	if err != nil {
		status = responses.FailureResponse("profile_id needs to be a number")
		resp, _ := json.Marshal(status)
		w.Write(resp)
		log.Println(status)
		return
	}

	err = database.RegisterPlayer(discordID, uint64(pid))

	if err != nil {
		status = responses.FailureResponse("FC cannot be found. Please connect to the limitless server and try again")
		resp, _ := json.Marshal(status)
		w.Write(resp)
		log.Println(status)
	} else {
		status = responses.SuccessResponse()
		status.Message = "User has been registered"
		resp, _ := json.Marshal(status)
		w.Write(resp)
		log.Println(status)
	}
}
