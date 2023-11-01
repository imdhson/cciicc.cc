package controller

import "net/http"

func HostHandler(w http.ResponseWriter, r *http.Request) {
	//incomingURL := r.URL.Path
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	a := []byte("HostTest")
	w.Write(a)
}
