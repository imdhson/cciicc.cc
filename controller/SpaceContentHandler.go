package controller

import (
	"html/template"
	"net/http"
	"ub/service"
	"ub/types"
)

type DataSpaceContent struct {
	Sp_id         string
	Sp_name       string
	Sp_view       int
	Sp_lastupdate string
	Sp_comments   []types.Sp_comment
}

func SpaceContentHandler(w http.ResponseWriter, r *http.Request, space_id string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/content.html")
	if err != nil {
		service.CriticalErr(err, "template html 로드")
	}

	//세션 가져오기
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		service.ErrHandler(getcookie_err, "SpaceContentHandler getCookie")
		http.Redirect(w, r, "/error", http.StatusFound)
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success || user.User_related_spaceid != space_id { //user success false일 때 or space_id 불일치시
		redirect_url := "/guest/" + space_id
		http.Redirect(w, r, redirect_url, http.StatusFound)
		return
	}

	space, getSpace_success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !getSpace_success {
		http.Redirect(w, r, "/error", http.StatusFound)
	}

	// 템플릿에 변수 설정
	data := DataSpaceContent{
		Sp_id:         space.Sp_id,
		Sp_name:       space.Sp_name,
		Sp_view:       space.Sp_view,
		Sp_lastupdate: space.Sp_lastupdate.String(),
		Sp_comments:   space.Sp_comments,
	}
	tmpl.Execute(w, data)
}
