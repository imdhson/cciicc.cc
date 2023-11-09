package controller

import (
	"net/http"
	"ub/service"
	"ub/types"
)

func SpaceHandler(w http.ResponseWriter, r *http.Request, space_id string) {
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		redirect_url := "/guest/" + space_id
		http.Redirect(w, r, redirect_url, http.StatusFound)
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}
	if getcookie_err == nil && user_success {
		redirect_url := "/space/" + user.User_related_spaceid
		http.Redirect(w, r, redirect_url, http.StatusFound)
	} else {
		redirect_url := "/guest/" + space_id
		http.Redirect(w, r, redirect_url, http.StatusFound)
	}

}
