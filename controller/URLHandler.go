package controller

import (
	"log"
	"net/http"
	"strings"
)

func URLHandler(w http.ResponseWriter, r *http.Request) {
	now_url_sliced := strings.Split(strings.ToLower(r.URL.Path), "/")
	now_url_sliced = now_url_sliced[1:]
	log.Println(now_url_sliced)

	switch now_url_sliced[0] {
	case "":
		MainHandler(w, r)
	case "host":
		HostHandler(w, r)
	case "guest":
		GuestHandler(w, r, now_url_sliced[1])
	case "error":
		ErrorPageHandler(w, r)
	case "assets":
		AssetsHanlder(w, r, r.URL.Path)
	case "space":
		switch now_url_sliced[1] {
		case "create":
			PostHandler_create_space(w, r)
		}
	default:
		ErrorPageHandler(w, r)
	}
}
