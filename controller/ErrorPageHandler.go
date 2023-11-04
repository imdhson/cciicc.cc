package controller

import (
	"net/http"
	"os"
	"ub/service"
)

func ErrorPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("wwwfiles/assets/err/index.html")
	service.CriticalErr(err)
	w.Write(wwwfile)
}
