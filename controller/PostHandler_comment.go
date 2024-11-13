package controller

import (
	"net/http"

	"cciicc/service"
	"cciicc/types"
)

func PostHandler_comment(w http.ResponseWriter, r *http.Request) {
	form_comment := r.FormValue("comment")

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
	comment := &types.Sp_chats{
		Sp_c_content:   form_comment,
		Sp_c_guestname: user.User_name,
	}

	service.AddCommentFrom_space_id(user.User_related_spaceid, comment)
	http.Redirect(w, r, "/space", http.StatusFound)
}
