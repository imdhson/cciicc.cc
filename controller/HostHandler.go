package controller

import (
	"html/template"
	"net/http"

	"cciicc.cc/service"
	"cciicc.cc/types"
)

type DataHost struct {
}

func HostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/host.html")
	if err != nil {
		service.CriticalErr(err, "template html 로드")
	}

	// 템플릿에 변수 설정
	type Data struct {
		Service_name         string
		Url_address          string
		Host_detail          string
		Host_spacename       string
		Host_username        string
		Host_spacename_input string
		Host_username_input  string
		Host_form_button     string

		Footer_terms string
	}
	data := Data{
		Service_name:         types.SERVICE_NAME,
		Url_address:          types.URL_ADDESS,
		Host_detail:          types.HOST_DETAIL,
		Host_spacename:       types.HOST_SPACENAME,
		Host_username:        types.HOST_USERNAME,
		Host_spacename_input: types.HOST_SPACENAME_INPUT,
		Host_username_input:  types.HOST_USERNAME_INPUT,
		Host_form_button:     types.HOST_FORM_BUTTON,
		Footer_terms:         types.FOOTER_TERMS,
	}
	tmpl.Execute(w, data)

}
