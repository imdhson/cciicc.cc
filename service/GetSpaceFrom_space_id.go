package service

import (
	"cciicc/types"
)

func GetSpaceFrom_space_id(space_id string) (*types.Space, bool) {
	spaces := types.GetInstance_spaces()
	var rst *types.Space
	rst_success := false
	for i, _ := range *spaces {
		if space_id == (*spaces)[i].Sp_id {
			rst_success = true
			rst = &(*spaces)[i]
			return rst, rst_success
		}
	}
	return rst, rst_success
}
