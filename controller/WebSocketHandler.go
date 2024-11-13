package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"cciicc/service"
	"cciicc/types"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	var user types.User
	var user_success bool
	session, getcookie_err := r.Cookie("ub_session")
	if getcookie_err != nil {
		service.ErrHandler(getcookie_err, "getcookie_err")
		w.WriteHeader(http.StatusForbidden)
		return
	} else { //오류 없을 때
		user, user_success = service.GetUserFromSession(session.Value)
	}

	if !user_success { //user success false일 때
		w.WriteHeader(http.StatusForbidden)
		return
	}

	space, _ := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	space_encoded, _ := json.MarshalIndent(space, " ", "	")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, space_encoded)
	if err != nil {
		log.Println("write:", err)
		return
	}

}
