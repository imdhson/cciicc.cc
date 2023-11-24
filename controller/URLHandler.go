package controller

import (
	"log"
	"net/http"
	"strings"

	"cciicc/service"
)

func URLHandler(w http.ResponseWriter, r *http.Request) {
	now_url_sliced := strings.Split(strings.ToLower(r.URL.Path), "/")
	now_url_sliced = append(now_url_sliced, "", "") //인덱스초과 방지용
	now_url_sliced = now_url_sliced[1:]

	switch now_url_sliced[0] {
	case "":
		log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
		MainHandler(w, r)
	case "host":
		log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
		HostHandler(w, r)
	case "guest":
		log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
		GuestHandler(w, r, now_url_sliced[1])
	case "error":
		log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
		ErrorPageHandler(w, r)
	case "assets":
		AssetsHanlder(w, r, r.URL.Path)
	case "space":
		switch now_url_sliced[1] {
		case "":
			SpaceHandler(w, r, now_url_sliced[1])
		case "create":
			log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
			PostHandler_create_space(w, r)
		case "join":
			log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
			PostHandler_join_space(w, r)
		case "json": //space/json
			SpaceJSONHandler(w, r)
		case "addcomment":
			log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
			PostHandler_comment(w, r)
		case "comment_rate":
			log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
			PostHandler_comment_rate(w, r)
		default:
			// space/[space_id] , space_id는 변조 위험이 있음으로 실제 핸들링시 확인 필요. UI로서의 space_id임.
			log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
			SpaceContentHandler(w, r, now_url_sliced[1])
		}
	default:
		log.Printf("%v/%v", service.GetIP(r), now_url_sliced)
		GuestHandler(w, r, now_url_sliced[0]) // url/guest/id 와 같으나 간소화된 url도 지원함
	}
}
