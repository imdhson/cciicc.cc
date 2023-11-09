package controller

import (
	"net/http"
	"time"
	"ub/service"
	"ub/types"
)

func PostHandler_comments(w http.ResponseWriter, r *http.Request) {

}
func PostHandler_create_space(w http.ResponseWriter, r *http.Request) {
	form_space_name := r.FormValue("spaceName")
	form_user_name := r.FormValue("userName")
	space_id := service.Create_space(form_space_name)
	if space_id == "" {
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	//랜덤 세션키 생성(space_id 기반)
	sessionkeyStr := service.Random_sessionkey_generator(space_id)

	session := http.Cookie{
		Name:     "ub_session",
		Value:    sessionkeyStr,
		Expires:  time.Now().Add(time.Hour * types.SESSION_EXPIRE_DURATION_HOURS),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &session)

	user := types.User{
		User_name:            form_user_name,
		User_sessionkey:      sessionkeyStr,
		User_isHost:          true,
		User_related_spaceid: space_id,
	}
	regist_success := service.UserRegist(user)
	if regist_success {
		http.Redirect(w, r, "/space", http.StatusFound)
	} else {
		http.Redirect(w, r, "/error", http.StatusFound)
	}

}

func PostHandler_join_space(w http.ResponseWriter, r *http.Request) {
	form_space_id := r.FormValue("spaceId")
	form_user_name := r.FormValue("userName")

	_, form_space_id_valid := service.GetSpaceFrom_space_id(form_space_id)
	if form_user_name == "" || !form_space_id_valid { //username이 없거나, spaceid가 유효하지 아니하면 에러
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	//랜덤 세션키 생성(space_id 기반)
	sessionkeyStr := service.Random_sessionkey_generator(form_space_id)

	session := http.Cookie{
		Name:     "ub_session",
		Value:    sessionkeyStr,
		Expires:  time.Now().Add(time.Hour * types.SESSION_EXPIRE_DURATION_HOURS),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &session)

	user := types.User{
		User_name:            form_user_name,
		User_sessionkey:      sessionkeyStr,
		User_isHost:          true,
		User_related_spaceid: form_space_id,
	}

	regist_success := service.UserRegist(user)
	if regist_success {
		http.Redirect(w, r, "/space", http.StatusFound)
	} else {
		http.Redirect(w, r, "/error", http.StatusFound)
	}

}
