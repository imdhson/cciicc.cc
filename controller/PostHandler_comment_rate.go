package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"cciicc.cc/service"
	"cciicc.cc/types"
)

func PostHandler_comment_rate(w http.ResponseWriter, r *http.Request) {
	//post 요청 수신
	form_like := r.FormValue("like_dislike")
	var form_like_bool bool
	if form_like == "like" {
		form_like_bool = true
	} else if form_like == "dislike" {
		form_like_bool = false
	} else {
		return
	}
	comment_c_id, err := strconv.Atoi(r.FormValue("comment_c_id"))
	service.ErrHandler(err, "comment_c_id : rate handling")

	fmt.Println(form_like, comment_c_id)

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

	service.RateCommentFrom_space_id(user.User_related_spaceid, &form_like_bool, &comment_c_id)

}
