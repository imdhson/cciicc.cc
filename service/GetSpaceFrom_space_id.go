package service

import "ub/types"

func GetSpaceFrom_space_id(space_id string) (types.Space, bool) {
	spaces := types.GetInstance_spaces()
	var rst types.Space
	rst_success := false
	for _, v := range *spaces {
		if space_id == v.Sp_id {
			rst_success = true
			rst = v
		}
	}
	return rst, rst_success
}
