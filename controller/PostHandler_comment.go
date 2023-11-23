package controller

import (
	"net/http"
	"strconv"
	"ub/service"
	"ub/types"
)

func PostHandler_comment(w http.ResponseWriter, r *http.Request) {
	form_comment := r.FormValue("comment")
	form_color := r.FormValue("color")
	form_color_int, form_color_int_err := strconv.Atoi(form_color)
	if form_color_int_err != nil {
		form_color_int = -1
	}

	session, getcookie_err := r.Cookie("ub_session")

	var user types.User
	var user_success bool
	if getcookie_err != nil {
		http.Redirect(w, r, "/error", http.StatusFound)
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}
	if getcookie_err != nil && !user_success {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	//comment 구조체 생성
	comment := &types.Sp_comment{
		Sp_c_content:   form_comment,
		Sp_c_guestname: user.User_name,
		Sp_c_color:     types.Sp_c_color(form_color_int),
	}

	service.AddCommentFrom_space_id(user.User_related_spaceid, comment)
	http.Redirect(w, r, "/space", http.StatusFound)
}
