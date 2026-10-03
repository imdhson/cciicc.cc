package service

import (
	"cciicc/types"
)

func GetSpaceFrom_space_id(space_id string) (*types.Space, bool) {
	spaces := *types.GetInstance_spaces()
	// Optimization: Iterate backwards because new spaces are appended
	// to the end of the slice, making them faster to find.
	for i := len(spaces) - 1; i >= 0; i-- {
		if space_id == spaces[i].Sp_id {
			return &spaces[i], true
		}
	}
	return nil, false
}
