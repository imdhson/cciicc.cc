package service

import (
	"fmt"
	"ub/types"
)

func Regist_space() string {
	spaces := types.GetInstance()
	space_id := Random_space_id_generator()
	valid := false
	for !valid { //혹시나 같은 것을 찾으면 다시 랜덤 돌리기위함
		for _, v := range *spaces { //순회하며 아이디같은지 찾고, 찾으면 랜덤
			if space_id == v.Sp_id {
				space_id = Random_space_id_generator()
			}
		}
		valid = true
	}

	space := types.Space{Sp_id: space_id,
		Sp_name: "no name", Sp_view: 0}
	*spaces = append(*spaces, space)

	for _, v := range *spaces { //스페이스공간 디버그용
		fmt.Println(v.Sp_id + "|" + v.Sp_name)
	}

	return space_id
}
