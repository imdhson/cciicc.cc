package controller

import (
	"cciicc/service"
	"cciicc/types"
	"encoding/json"
	"net/http"
	"strconv"
)

func PostHandler_file_context(w http.ResponseWriter, r *http.Request) {
	pageNum := r.FormValue("file_context")
	session, getcookie_err := r.Cookie("ub_session")

	var user types.User
	var user_success bool
	if getcookie_err != nil {
		http.Error(w, "user is not host", http.StatusForbidden)
		return
	} else {
		user, user_success = service.GetUserFromSession(session.Value)
	}
	if getcookie_err != nil && !user_success {
		http.Error(w, "user is not host", http.StatusForbidden)
		return
	}
	if !user.User_isHost {
		http.Error(w, "user is not host", http.StatusForbidden)
		return
	}
	space, space_success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !space_success {
		http.Error(w, "cannot find space", http.StatusInternalServerError)
	}
	//space에 업데이트
	pageNum_i, err := strconv.Atoi(pageNum)
	if err != nil {
		service.ErrHandler(err, "strconv posthandler file context")
		http.Error(w, "strconv post atoi error at file context", http.StatusBadRequest)
		return
	}
	space.Sp_file_context = pageNum_i

	//같은 ws_space에 websocket broadcast 시도
	ws_hub := GetInstance_ws_hub()
	ws_space := ws_hub.Ws_GetOrCreateSpace(user.User_related_spaceid)
	ws_file_context := types.New_Sp_ws_type_file_context(space.Sp_file_status, pageNum)
	ws_file_context_encoded, err := json.MarshalIndent(ws_file_context, " ", "	")
	service.ErrHandler(err, "posthandler file context json")
	ws_hub.broadcast([]byte(ws_file_context_encoded), ws_space)

	http.Redirect(w, r, "/space", http.StatusFound)
}
