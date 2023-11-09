package controller

import (
	"log"
	"net/http"
	"strings"
)

func URLHandler(w http.ResponseWriter, r *http.Request) {
	now_url_sliced := strings.Split(strings.ToLower(r.URL.Path), "/")
	now_url_sliced = append(now_url_sliced, "", "") //인덱스초과 방지용
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
		case "":
			SpaceHandler(w, r, now_url_sliced[1])
		case "create":
			PostHandler_create_space(w, r)
		case "join":
			PostHandler_join_space(w, r)
		case "json": //space/json
			SpaceJSONHandler(w, r)
		default:
			// space/[space_id] , space_id는 변조 위험이 있음으로 실제 핸들링시 확인 필요. UI로서의 space_id임.
			SpaceContentHandler(w, r, now_url_sliced[1]) //수정 필요
		}
	default:
		ErrorPageHandler(w, r)
	}
}
