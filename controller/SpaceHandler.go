package controller

import "net/http"

func SpaceHandler(w http.ResponseWriter, r *http.Request) {
	r.Cookie("ub_session")
}
