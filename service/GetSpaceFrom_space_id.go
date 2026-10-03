package service

import (
	"cciicc/types"
)

func GetSpaceFrom_space_id(space_id string) (*types.Space, bool) {
	spaces := *types.GetInstance_spaces()
	for i := 0; i < len(spaces); i++ {
		if space_id == spaces[i].Sp_id {
			return &spaces[i], true
		}
	}
	return nil, false
}
