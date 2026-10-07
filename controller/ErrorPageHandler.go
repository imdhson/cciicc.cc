package controller

import (
	"net/http"

	"cciicc/types"
)

func ErrorPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 로드 (캐시 사용)
	tmpl := GetTmplError()
	type Data struct {
		Service_name  string
		Error_title   string
		Error_content string
		Error_main    string
		Footer_terms  string
	}

	// 템플릿에 변수 설정
	data := Data{
		Service_name:  types.SERVICE_NAME,
		Error_title:   types.ERROR_TITLE,
		Error_content: types.ERROR_CONTENT,
		Error_main:    types.ERROR_MAIN,
		Footer_terms:  types.FOOTER_TERMS,
	}
	tmpl.Execute(w, data)
}
