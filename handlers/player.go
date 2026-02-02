package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MKW-Limitless-team/limitless-types/json_response"
	"github.com/nwoik/Limitless-API/database"
)

func Player(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	discordID := r.URL.Query().Get("discord_id")
	player, user, err := database.GetPlayerInfo(discordID)

	if err != nil {
		resp, _ := json.Marshal(json_response.FailureResponse("Failed to get player"))
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(json_response.SuccessResponse(player, user))
	w.Write(resp)
}

func Register(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	discordID := r.URL.Query().Get("discord_id")
	profileID := r.URL.Query().Get("profile_id")

	pid, err := strconv.Atoi(profileID)

	if err != nil {
		resp, _ := json.Marshal(json_response.FailureResponse("profile_id needs to be a number"))
		w.Write(resp)
		return
	}

	err = database.RegisterPlayer(discordID, uint64(pid))

	if err != nil {
		resp, _ := json.Marshal(json_response.FailureResponse("FC cannot be found. Please connect to the limitless server and try again"))
		w.Write(resp)
	} else {
		resp, _ := json.Marshal(json_response.SuccessResponse("User has been registered"))
		w.Write(resp)
	}
}
