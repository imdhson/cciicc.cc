package controller

import (
	"html/template"
	"log"
	"net/http"
)

type DataGuest struct {
}

func GuestHandler(w http.ResponseWriter, r *http.Request, space_id string) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/guest.html")
	if err != nil {
		log.Fatal(err)
	}

	// 템플릿에 변수 설정
	data := DataGuest{}
	tmpl.Execute(w, data)
}
