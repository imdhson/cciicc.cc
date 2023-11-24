package controller

import (
	"encoding/json"
	"net/http"

	"cciicc/types"

	"cciicc/service"
)

func SpaceJSONHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		service.ErrHandler(getcookie_err, "getcookie_err")
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	} else { //오류 없을 때
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success { //user success false일 때
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	space, _ := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	space_encoded, err := json.MarshalIndent(space, " ", "	")
	service.ErrHandler(err, "space_encoded")
	w.Write(space_encoded)
}
