package controller

import (
	"encoding/json"
	"net/http"

	"cciicc/service"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("ub_session")
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	user, success := service.GetUserFromSession(session.Value)
	if !success || !user.User_isHost {
		http.Error(w, "Forbidden: Only host can export data", http.StatusForbidden)
		return
	}

	space, success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !success {
		http.Error(w, "Space not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=\"space_record.json\"")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(space); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
