package controller

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
}

func HostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//space_id := service.Regist_space()
	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/register.html")
	if err != nil {
		log.Fatal(err)
	}

	// 템플릿에 변수 설정
	data := Data{}
	tmpl.Execute(w, data)

}
