package controller

import (
	"encoding/json"
	"net/http"

	"cciicc/service"
	"cciicc/types"
)

func PostHandler_emoji(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, err := r.Cookie("ub_session")
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	user, success := service.GetUserFromSession(session.Value)
	if !success {
		http.Error(w, "Forbidden", http.StatusForbidden)
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
