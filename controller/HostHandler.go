package controller

import (
	"net/http"
	"ub/service"
)

func HostHandler(w http.ResponseWriter, r *http.Request) {
	space_id := service.Regist_space()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write()
}
