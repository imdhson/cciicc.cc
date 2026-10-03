package controller

import (
	"cciicc/service"
	"cciicc/types"
	"encoding/json"
	"net/http"
)

func PostHandler_emoji(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, getcookie_err := r.Cookie("ub_session")

	var user types.User
	var user_success bool
	if getcookie_err != nil {
		http.Error(w, "Invalid session", http.StatusForbidden)
		return
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}
	if !user_success {
		http.Error(w, "Invalid session", http.StatusForbidden)
		return
	}

	emoji := r.FormValue("emoji")
	if emoji == "" {
		http.Error(w, "Emoji is required", http.StatusBadRequest)
		return
	}

	ws_hub := types.GetInstance_ws_hub()
	ws_space := ws_hub.Ws_GetOrCreateSpace(user.User_related_spaceid)

	emoji_payload := types.New_Sp_ws_type_emoji(emoji)
	emoji_json, err := json.Marshal(emoji_payload)
	if err != nil {
		service.ErrHandler(err, "emoji json marshal error")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	ws_hub.Broadcast(emoji_json, ws_space)

	w.WriteHeader(http.StatusOK)
}
