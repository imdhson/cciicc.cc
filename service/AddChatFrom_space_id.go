package service

import (
	"time"

	"cciicc/types"
)

func AddChatFrom_space_id(space_id string, comment *types.Sp_chat) bool {
	spacesPtr := types.GetInstance_spaces()
	spaces := *spacesPtr
	tmpi := -1
	for i := len(spaces) - 1; i >= 0; i-- {
		if spaces[i].Sp_id == space_id {
			tmpi = i
			break
		}
	}

	if tmpi == -1 {
		return false
	}

	comment.Sp_c_id = len(spaces[tmpi].Sp_chats)
	spaces[tmpi].Sp_chats = append(spaces[tmpi].Sp_chats, *comment)
	spaces[tmpi].Sp_lastupdate = time.Now()
	*spacesPtr = spaces
	return true
}
