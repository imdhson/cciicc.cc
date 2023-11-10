package service

import "ub/types"

func AddCommentFrom_space_id(space_id string, comment *types.Sp_comment) {
	spaces := types.GetInstance_spaces()
	var tmpi int
	for i, v := range *spaces {
		if v.Sp_id == space_id {
			tmpi = i
			break
		}
	}
	(*spaces)[tmpi].Sp_comments = append((*spaces)[tmpi].Sp_comments, *comment)
}
