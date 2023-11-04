package controller

import "net/http"

func GuestHandler(w http.ResponseWriter, r *http.Request, spaceId string) {
	//incomingURL := r.URL.Path
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	a := []byte("GuestTest")
	w.Write(a)
}
