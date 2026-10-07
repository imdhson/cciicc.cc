package controller

import (
	"net/http"

	"cciicc/service"
	"cciicc/types"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 로드 (캐시 사용)
	tmpl := GetTmplMain()

	// 템플릿에 변수 설정
	type Data struct {
		Service_name    string
		Service_detail  string
		Footer_terms    string
		Main_next_up    string
		Main_if_next_up bool //상수로 불러오지 아니함
		Main_host       string
		Main_guest      string
	}
	data := Data{
		Service_name:    types.SERVICE_NAME,
		Service_detail:  types.SERVICE_DETAIL,
		Footer_terms:    types.FOOTER_TERMS,
		Main_next_up:    types.MAIN_NEXT_UP,
		Main_if_next_up: false,
		Main_host:       types.MAIN_HOST,
		Main_guest:      types.MAIN_GUEST,
	}

	//이어서 시작하기 버튼 노출 여부 결정
	session, getcookie_err := r.Cookie("ub_session")
	var user_success bool
	if getcookie_err == nil {
		_, user_success = service.GetUserFromSession(session.Value)
	}
	if user_success {
		data.Main_if_next_up = true
	}
	tmpl.Execute(w, data)
}
