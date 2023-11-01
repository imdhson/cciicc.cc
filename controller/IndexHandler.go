package controller

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//incomingURL := r.URL.Path
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	a := []byte("IndexTest")
	w.Write(a)
}
