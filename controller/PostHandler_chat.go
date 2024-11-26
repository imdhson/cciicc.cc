package controller

import (
	"encoding/json"
	"net/http"

	"cciicc/service"
	"cciicc/types"
)

func PostHandler_chat(w http.ResponseWriter, r *http.Request) {
	form_chat := r.FormValue("chat")

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

	//chat 구조체 생성
	sp_chat := &types.Sp_chat{
		Sp_c_content:   form_chat,
		Sp_c_guestname: user.User_name,
	}
	service.AddChatFrom_space_id(user.User_related_spaceid, sp_chat)

	//추가하고 같은 ws_space에 websocket broadcast 시도
	ws_hub := types.GetInstance_ws_hub()
	ws_space := ws_hub.Ws_GetOrCreateSpace(user.User_related_spaceid)
	ws_chat_content := types.New_Sp_ws_type_chat(sp_chat.Sp_c_guestname, sp_chat.Sp_c_content)
	ws_chat_content_encoded, err := json.MarshalIndent(ws_chat_content, " ", "	")
	service.ErrHandler(err, "posthandler chat json")
	ws_hub.Broadcast([]byte(ws_chat_content_encoded), ws_space)

	http.Redirect(w, r, "/space", http.StatusFound)
}
