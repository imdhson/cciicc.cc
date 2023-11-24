package service

import (
	"time"

	"cciicc/types"
)

func RateCommentFrom_space_id(space_id string, rate_like *bool, comment_c_id *int) {
	spaces := types.GetInstance_spaces()
	var tmpi_space int
	var tmpi_comment int
	for i, v := range *spaces {
		if v.Sp_id == space_id {
			tmpi_space = i //space 찾아서 tmpi 에 저장
			break
		}
	}
	for i, v := range (*spaces)[tmpi_space].Sp_comments {
		if v.Sp_c_id == *comment_c_id {
			tmpi_comment = i //space안 comments에서 comment객체 찾아서 인덱스 저장
		}
	}
	if *rate_like {
		(*spaces)[tmpi_space].Sp_comments[tmpi_comment].Sp_c_rate += 1
	} else {
		(*spaces)[tmpi_space].Sp_comments[tmpi_comment].Sp_c_rate -= 1
	}
	(*spaces)[tmpi_space].Sp_lastupdate = time.Now()

}
