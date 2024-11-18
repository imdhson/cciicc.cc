package controller

import (
	"html/template"
	"net/http"

	"cciicc/types"

	"cciicc/service"
)

func SpaceContentHandler(w http.ResponseWriter, r *http.Request, space_id string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//세션 가져오기
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		redirect_url := "/guest/" + space_id
		http.Redirect(w, r, redirect_url, http.StatusFound)
		return
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success || user.User_related_spaceid != space_id { //user success false일 때 or space_id 불일치시
		redirect_url := "/guest/" + space_id
		http.Redirect(w, r, redirect_url, http.StatusFound)
		return
	}

	var tmpl *template.Template
	if user.User_isHost { // user가 host이면 spacecontent host 템플릿 반환
		// 템플릿 파일 로드
		tmpl_i, err := template.ParseFiles("wwwfiles/content_host.html")
		tmpl = tmpl_i
		if err != nil {
			service.CriticalErr(err, "template html 로드 host")
		}
	} else { // user가 guest이면 space content guest 템플릿 반환
		tmpl_i, err := template.ParseFiles("wwwfiles/content_guest.html")
		tmpl = tmpl_i
		if err != nil {
			service.CriticalErr(err, "template html 로드 guest")
		}
	}

	space, getSpace_success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !getSpace_success {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	//space에 조회수 1 추가
	space.Sp_view += 1
	// 템플릿에 변수 설정
	type DataSpaceContent struct {
		Service_name  string
		Url_address   string
		Sp_id         string
		Sp_name       string
		Sp_view       int
		Sp_lastupdate string
		Sp_chats      []types.Sp_chat
		Sp_png_path   string
		User_name     string

		Content_order      string
		Content_view_count string
		Content_send       string

		Footer_terms string
	}
	data := DataSpaceContent{
		Service_name:  types.SERVICE_NAME,
		Url_address:   types.URL_ADDESS,
		User_name:     user.User_name,
		Sp_id:         space.Sp_id,
		Sp_name:       space.Sp_name,
		Sp_view:       space.Sp_view,
		Sp_lastupdate: space.Sp_lastupdate.String(),
		Sp_chats:      space.Sp_chats,
		Sp_png_path:   types.URL_ADDESS + "/assets/space_qr/" + space_id + ".png",

		Content_view_count: types.CONTENT_VIEW_COUNT,
		Content_order:      types.CONTENT_ORDER,
		Content_send:       types.CONTENT_SEND,

		Footer_terms: types.FOOTER_TERMS,
	}
	tmpl.Execute(w, data)
}
