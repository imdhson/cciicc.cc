package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cciicc/service"
	"cciicc/types"
)

func PostHandler_media_sync(w http.ResponseWriter, r *http.Request) {
	currentTimeStr := r.FormValue("currentTime")
	isPausedStr := r.FormValue("isPaused")

	session, getcookie_err := r.Cookie("ub_session")

	var user types.User
	var user_success bool
	if getcookie_err != nil {
		http.Error(w, "user is not host", http.StatusForbidden)
		return
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success || !user.User_isHost {
		http.Error(w, "user is not host", http.StatusForbidden)
		return
	}

	currentTime, err := strconv.ParseFloat(currentTimeStr, 64)
	if err != nil {
		http.Error(w, "invalid currentTime", http.StatusBadRequest)
		return
	}
	isPaused, err := strconv.ParseBool(isPausedStr)
	if err != nil {
		http.Error(w, "invalid isPaused", http.StatusBadRequest)
		return
	}

	ws_hub := types.GetInstance_ws_hub()
	ws_space := ws_hub.Ws_GetOrCreateSpace(user.User_related_spaceid)

	ws_media_sync := types.New_Sp_ws_type_media_sync(currentTime, isPaused)
	ws_media_sync_encoded, err := json.MarshalIndent(ws_media_sync, " ", "\t")
	if err != nil {
		service.ErrHandler(err, "posthandler media sync json")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	ws_hub.Broadcast([]byte(ws_media_sync_encoded), ws_space)

	w.WriteHeader(http.StatusOK)
}
