package controller

import (
	"crypto/md5"
	"fmt"
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
		ErrorPageHandler(w, r)
	}
	//랜덤 세션키 생성(space_id 기반)
	rand_key := service.Random_sessionkey_generator(space_id)
	//해쉬함수
	hash := md5.New()
	hash.Write([]byte(rand_key))
	hashSum := hash.Sum(nil)
	sessionkeyStr := fmt.Sprintf("%x", hashSum)

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
		http.Redirect(w, r, "/space", http.StatusFound) //링크 수정 필요
	} else {
		http.Redirect(w, r, "/error", http.StatusFound)
	}

}
