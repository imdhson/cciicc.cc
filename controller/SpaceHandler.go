package controller

import (
	"fmt"
	"net/http"
	"ub/service"
	"ub/types"
)

func SpaceHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		service.ErrHandler(getcookie_err)
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}
	if getcookie_err == nil && user_success {
		fmt.Println(user)
		redirect_url := "/space/" + user.User_related_spaceid
		http.Redirect(w, r, redirect_url, http.StatusFound)
	} else {
		http.Redirect(w, r, "/error", http.StatusFound)
	}

}
