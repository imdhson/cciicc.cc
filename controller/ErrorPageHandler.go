package controller

import (
	"net/http"
	"os"
	"ub/service"
)

func ErrorPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	wwwfile, err := os.ReadFile("wwwfiles/assets/error/index.html")
	service.ErrHandler(err, "ErrorPageHandler")
	w.Write(wwwfile)
}
