package controller

import (
	"html/template"
	"net/http"

	"cciicc/service"
	"cciicc/types"

	"github.com/gorilla/csrf"
)

func GuestHandler(w http.ResponseWriter, r *http.Request, space_id string) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 로드 (캐시 사용)
	tmpl := GetTmplGuest()
	type Data struct {
		Sp_id                string
		Service_name         string
		Guest_detail         string
		Guest_spaceid        string
		Guest_spaceid_input  string
		Guest_username       string
		Guest_username_input string
		Guest_form_button    string
		Footer_terms         string
		CsrfField            template.HTML
	}

	// 템플릿에 변수 설정
	data := Data{
		Sp_id: space_id,

		Service_name: types.SERVICE_NAME,

		Guest_detail:         types.GUEST_DETAIL,
		Guest_spaceid:        types.GUEST_SPACEID,
		Guest_spaceid_input:  types.GUEST_SPACEID_INPUT + service.Random_space_id_generator(),
		Guest_username:       types.GUEST_USERNAME,
		Guest_username_input: types.GUEST_USERNAME_INPUT,
		Guest_form_button:    types.GUEST_FORM_BUTTON,

		Footer_terms: types.FOOTER_TERMS,
		CsrfField:    csrf.TemplateField(r),
	}
	tmpl.Execute(w, data)
}
