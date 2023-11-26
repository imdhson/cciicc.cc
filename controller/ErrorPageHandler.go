package controller

import (
	"html/template"
	"net/http"

	"cciicc/service"
	"cciicc/types"
)

func ErrorPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/assets/error/index.html")
	if err != nil {
		service.CriticalErr(err, "template html 로드")
	}
	type Data struct {
		Service_name  string
		Error_title   string
		Error_content string
		Error_main    string
		Url_address   string
		Footer_terms  string
	}

	// 템플릿에 변수 설정
	data := Data{
		Service_name:  types.SERVICE_NAME,
		Url_address:   types.URL_ADDESS,
		Error_title:   types.ERROR_TITLE,
		Error_content: types.ERROR_CONTENT,
		Error_main:    types.ERROR_MAIN,
		Footer_terms:  types.FOOTER_TERMS,
	}
	tmpl.Execute(w, data)
}
