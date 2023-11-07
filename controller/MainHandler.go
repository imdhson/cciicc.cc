package controller

import (
	"net/http"
	"os"
	"ub/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("wwwfiles/main.html")
	service.ErrHandler(err, "os.ReadFile('wwwfiles/main.html')")
	w.Write(wwwfile)
}
