package controller

import (
	"html/template"
	"net/http"
	"ub/service"
)

type DataHost struct {
}

func HostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//space_id := service.Regist_space()
	// 템플릿 파일 로드
	tmpl, err := template.ParseFiles("wwwfiles/host.html")
	if err != nil {
		service.CriticalErr(err, "template html 로드")
	}

	// 템플릿에 변수 설정
	data := DataHost{}
	tmpl.Execute(w, data)

}
