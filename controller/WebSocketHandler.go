package controller

import (
	"net/http"

	"cciicc/service"
	"cciicc/types"

	"github.com/gorilla/websocket"
)

func WebSocketHandler(h *types.Ws_Hub, w http.ResponseWriter, r *http.Request) {
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

	// WebSocket 연결을 위한 업그레이더 설정
	var upgrader = websocket.Upgrader{} // ReadBufferSize:  1024,WriteBufferSize: 1024, 주석 처리 시 기본 옵션 사용

	conn, err := upgrader.Upgrade(w, r, nil)
	service.ErrHandler(err, "websocket upgrader")

	space_content, space_content_success := service.GetSpaceFrom_space_id(user.User_related_spaceid)
	if !space_content_success {
		return
	}
	// space_content_encoded, err := json.MarshalIndent(space_content, " ", "	")
	// service.ErrHandler(err, "wssockethandler jsonmarshal")
	conn.WriteJSON(space_content)

	// 새 클라이언트를 생성하고 해당 게시글에 추가합니다.
	ws_client := &types.Ws_Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	ws_space := h.Ws_GetOrCreateSpace(user.User_related_spaceid) //space_id 입력으로 ws_space를 얻어줌

	ws_space.Mu.Lock()
	ws_space.Ws_clients[ws_client] = true
	ws_space.Mu.Unlock()

	// 클라이언트의 읽기와 쓰기를 처리하는 고루틴을 시작합니다.
	// go h.readPump(ws_client, ws_space)
	go h.WritePump(ws_client, ws_space)
}
