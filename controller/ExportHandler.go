package controller

import (
	"cciicc/service"
	"cciicc/types"
	"encoding/json"
	"net/http"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
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

	if !user.User_isHost {
		http.Error(w, "user is not host", http.StatusForbidden)
		return
	}

	space, success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !success {
		http.Error(w, "Space not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=\"space_record.json\"")

	json.NewEncoder(w).Encode(space)
}
