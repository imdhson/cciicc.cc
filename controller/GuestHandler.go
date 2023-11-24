package controller

import (
	"html/template"
	"net/http"

	"cciicc/service"
	"cciicc/types"
)

func GuestHandler(w http.ResponseWriter, r *http.Request, space_id string) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/guest.html")
	if err != nil {
		service.CriticalErr(err, "template html 로드")
	}
	type Data struct {
		Sp_id                string
		Url_address          string
		Service_name         string
		Guest_detail         string
		Guest_spaceid        string
		Guest_spaceid_input  string
		Guest_username       string
		Guest_username_input string
		Guest_form_button    string
	}

	// 템플릿에 변수 설정
	data := Data{
		Sp_id: space_id,

		Service_name: types.SERVICE_NAME,
		Url_address:  types.URL_ADDESS,

		Guest_detail:         types.GUEST_DETAIL,
		Guest_spaceid:        types.GUEST_SPACEID,
		Guest_spaceid_input:  types.GUEST_SPACEID_INPUT + service.Random_space_id_generator(),
		Guest_username:       types.GUEST_USERNAME,
		Guest_username_input: types.GUEST_USERNAME_INPUT,
		Guest_form_button:    types.GUEST_FORM_BUTTON,
	}
	tmpl.Execute(w, data)
}
