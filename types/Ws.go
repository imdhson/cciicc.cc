package types

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var ws_hub *Ws_Hub //single ton

// Ws_Client는 개별 WebSocket 연결을 나타냅니다.
type Ws_Client struct {
	Conn *websocket.Conn // WebSocket 연결
	Send chan []byte     // 클라이언트로 보낼 메시지를 담는 채널
}

// Ws_Space는 하나의 게시글과 관련된 WebSocket 클라이언트들을 관리합니다.
type Ws_Space struct {
	Space_id   string              // 게시글 ID
	Ws_clients map[*Ws_Client]bool // 연결된 클라이언트 맵
	Mu         sync.Mutex          // 동시성 제어를 위한 뮤텍스
}

// Ws_Hub는 모든 게시글의 WebSocket 연결을 관리합니다.
type Ws_Hub struct {
	Ws_spaces map[string]*Ws_Space // 게시글 ID를 키로 하는 ws_space 맵
	Mu        sync.Mutex           // 동시성 제어를 위한 뮤텍스
}

func GetInstance_ws_hub() *Ws_Hub { //single ton
	if ws_hub == nil {
		ws_hub = &Ws_Hub{
			Ws_spaces: make(map[string]*Ws_Space),
		}
	}
	return ws_hub
}

// 게시글 ID에 해당하는 ws_space를 가져오거나 새로 생성합니다.
func (h *Ws_Hub) Ws_GetOrCreateSpace(space_id string) *Ws_Space {
	h.Mu.Lock()
	defer h.Mu.Unlock()

	if ws_space, exists := h.Ws_spaces[space_id]; exists {
		return ws_space
	}

	ws_space := &Ws_Space{
		Space_id:   space_id,
		Ws_clients: make(map[*Ws_Client]bool),
	}
	h.Ws_spaces[space_id] = ws_space
	return ws_space
}

// Ws_RemoveSpace는 특정 게시글 ID에 해당하는 ws_space를 삭제하고 관련 클라이언트 연결을 종료합니다.
func (h *Ws_Hub) Ws_RemoveSpace(space_id string) {
	h.Mu.Lock()
	ws_space, exists := h.Ws_spaces[space_id]
	if !exists {
		h.Mu.Unlock()
		return
	}
	delete(h.Ws_spaces, space_id)
	h.Mu.Unlock()

	// 게시글에 연결된 모든 클라이언트의 연결을 종료합니다.
	ws_space.Mu.Lock()
	for ws_client := range ws_space.Ws_clients {
		close(ws_client.Send)
		ws_client.Conn.Close()
		delete(ws_space.Ws_clients, ws_client)
	}
	ws_space.Mu.Unlock()

	log.Printf("웹소켓 spaceid: %s가 삭제되었습니다.", space_id)
}

// // 클라이언트로부터 메시지를 읽는 펌프 함수입니다.
// func (h *Ws_Hub) ReadPump(ws_client *Ws_Client, ws_space *Ws_Space) {
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
func (h *Ws_Hub) WritePump(ws_client *Ws_Client, ws_space *Ws_Space) {
	defer func() {
		ws_client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-ws_client.Send:
			if !ok {
				// 채널이 닫혔으면 연결을 종료합니다.
				ws_client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			// 메시지를 클라이언트로 전송합니다.
			w, err := ws_client.Conn.NextWriter(websocket.TextMessage)
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
func (h *Ws_Hub) Broadcast(message []byte, ws_space *Ws_Space) {
	ws_space.Mu.Lock()
	defer ws_space.Mu.Unlock()
	for ws_client := range ws_space.Ws_clients {
		select {
		case ws_client.Send <- message: //message를 전송함
		default:
			// 클라이언트가 메시지를 받을 수 없으면 연결을 종료합니다.
			log.Println("ws client close 됨")
			close(ws_client.Send)
			delete(ws_space.Ws_clients, ws_client)
		}
	}
}
