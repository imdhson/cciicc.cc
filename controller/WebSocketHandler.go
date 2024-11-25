package controller

import (
	"log"
	"net/http"
	"sync"

	"cciicc/service"
	"cciicc/types"

	"github.com/gorilla/websocket"
)

var ws_hub *Ws_Hub //single ton

// Ws_Client는 개별 WebSocket 연결을 나타냅니다.
type Ws_Client struct {
	conn *websocket.Conn // WebSocket 연결
	send chan []byte     // 클라이언트로 보낼 메시지를 담는 채널
}

// Ws_Space는 하나의 게시글과 관련된 WebSocket 클라이언트들을 관리합니다.
type Ws_Space struct {
	space_id   string              // 게시글 ID
	ws_clients map[*Ws_Client]bool // 연결된 클라이언트 맵
	mu         sync.Mutex          // 동시성 제어를 위한 뮤텍스
}

// Ws_Hub는 모든 게시글의 WebSocket 연결을 관리합니다.
type Ws_Hub struct {
	ws_spaces map[string]*Ws_Space // 게시글 ID를 키로 하는 ws_space 맵
	mu        sync.Mutex           // 동시성 제어를 위한 뮤텍스
}

func GetInstance_ws_hub() *Ws_Hub { //single ton
	if ws_hub == nil {
		ws_hub = &Ws_Hub{
			ws_spaces: make(map[string]*Ws_Space),
		}
	}
	return ws_hub
}

// 게시글 ID에 해당하는 ws_space를 가져오거나 새로 생성합니다.
func (h *Ws_Hub) Ws_GetOrCreateSpace(space_id string) *Ws_Space {
	h.mu.Lock()
	defer h.mu.Unlock()

	if ws_space, exists := h.ws_spaces[space_id]; exists {
		return ws_space
	}

	ws_space := &Ws_Space{
		space_id:   space_id,
		ws_clients: make(map[*Ws_Client]bool),
	}
	h.ws_spaces[space_id] = ws_space
	return ws_space
}

// Ws_RemoveSpace는 특정 게시글 ID에 해당하는 ws_space를 삭제하고 관련 클라이언트 연결을 종료합니다.
func (h *Ws_Hub) Ws_RemoveSpace(space_id string) {
	h.mu.Lock()
	ws_space, exists := h.ws_spaces[space_id]
	if !exists {
		h.mu.Unlock()
		return
	}
	delete(h.ws_spaces, space_id)
	h.mu.Unlock()

	// 게시글에 연결된 모든 클라이언트의 연결을 종료합니다.
	ws_space.mu.Lock()
	for ws_client := range ws_space.ws_clients {
		close(ws_client.send)
		ws_client.conn.Close()
		delete(ws_space.ws_clients, ws_client)
	}
	ws_space.mu.Unlock()

	log.Printf("웹소켓 spaceid: %s가 삭제되었습니다.", space_id)
}

// // 클라이언트로부터 메시지를 읽는 펌프 함수입니다.
// func (h *Ws_Hub) readPump(ws_client *Ws_Client, ws_space *Ws_Space) {
// 	defer func() {
// 		ws_space.mu.Lock()
// 		delete(ws_space.ws_clients, ws_client)
// 		ws_space.mu.Unlock()
// 		ws_client.conn.Close()
// 	}()

// 	for {
// 		// 클라이언트로부터 메시지를 읽습니다.
// 		_, message, err := ws_client.conn.ReadMessage()

// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("error: %v", err)
// 			}
// 			break
// 		}

// 		// //[]byte를 json 변환
// 		var jsonData types.Sp_ws_type_file_context
// 		err = json.Unmarshal(message, &jsonData)
// 		service.ErrHandler(err, "ws handler json unmarshal")
// 		// 받은 메시지를 같은 게시글의 모든 클라이언트에게 브로드캐스트합니다.
// 		// h.broadcast(message, ws_space)
// 	}
// }

// 클라이언트로 메시지를 보내는 펌프 함수입니다.
func (h *Ws_Hub) writePump(ws_client *Ws_Client, ws_space *Ws_Space) {
	defer func() {
		ws_client.conn.Close()
	}()

	for {
		select {
		case message, ok := <-ws_client.send:
			if !ok {
				// 채널이 닫혔으면 연결을 종료합니다.
				ws_client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 메시지를 클라이언트로 전송합니다.
			w, err := ws_client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// 메시지를 게시글의 모든 클라이언트에게 브로드캐스트합니다.
func (h *Ws_Hub) broadcast(message []byte, ws_space *Ws_Space) {
	ws_space.mu.Lock()
	defer ws_space.mu.Unlock()
	for ws_client := range ws_space.ws_clients {
		select {
		case ws_client.send <- message: //message를 전송함
		default:
			// 클라이언트가 메시지를 받을 수 없으면 연결을 종료합니다.
			log.Println("ws client close 됨")
			close(ws_client.send)
			delete(ws_space.ws_clients, ws_client)
		}
	}
}

func (h *Ws_Hub) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
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
	ws_client := &Ws_Client{
		conn: conn,
		send: make(chan []byte, 256),
	}
	ws_space := h.Ws_GetOrCreateSpace(user.User_related_spaceid) //space_id 입력으로 ws_space를 얻어줌

	ws_space.mu.Lock()
	ws_space.ws_clients[ws_client] = true
	ws_space.mu.Unlock()

	// 클라이언트의 읽기와 쓰기를 처리하는 고루틴을 시작합니다.
	// go h.readPump(ws_client, ws_space)
	go h.writePump(ws_client, ws_space)
}
