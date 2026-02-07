package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/MKW-Limitless-team/limitless-types/responses"
	"github.com/MKW-Limitless-team/limitless-types/wwfc"
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
	friendCode := r.URL.Query().Get("friend_code")

	fc, err := strconv.Atoi(strings.ReplaceAll(friendCode, "-", ""))

	if err != nil {
		status = responses.FailureResponse("Friend-code needs to be a number")
		resp, _ := json.Marshal(status)
		w.Write(resp)
		log.Println(status)
		return
	}

	_, _, err = database.GetPlayerInfo(discordID)

	if err == nil {
		status = responses.FailureResponse("User already registered")
		resp, _ := json.Marshal(status)
		w.Write(resp)
		log.Println(status)
		return
	}

	userIDs := database.GetProfileIDs()
	var profileID uint64
	for _, id := range userIDs {
		if uint64(fc) == uint64(wwfc.PidToFC(id)) {
			profileID = id
			break
		}
	}

	err = database.RegisterPlayer(discordID, profileID)

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
