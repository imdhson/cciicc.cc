package controller

import (
	"fmt"
	"net/http"
	"strings"
)

func URLHandler(w http.ResponseWriter, r *http.Request) {
	now_url_sliced := strings.Split(strings.ToLower(r.URL.Path), "/")
	now_url_sliced = now_url_sliced[1:]
	fmt.Println(now_url_sliced)

	switch now_url_sliced[0] {
	case "host":
		HostHandler(w, r)
	case "guest":
		GuestHandler(w, r, now_url_sliced[1])
	case "error":
		ErrorPageHandler(w, r)
	}
}
