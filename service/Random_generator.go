package service

import (
	"math/rand"
	"strconv"
	"ub/types"
)

func Random_space_id_generator() string {
	spaces := types.GetInstance_space()
	rand_int := rand.Intn(999)
	rand_char := string(rune(rand.Intn(26) + 97))
	space_id := rand_char + strconv.Itoa(rand_int)

	valid := false
	for !valid { //혹시나 같은 것을 찾으면 다시 랜덤 돌리기위함
		for _, v := range *spaces { //순회하며 아이디같은지 찾고, 찾으면 랜덤
			if space_id == v.Sp_id {
				space_id = Random_space_id_generator()
			}
		}
		valid = true
	}
	return space_id
}

func Random_sessionkey_generator(space_id string) string {
	users := types.GetInstance_users()
	rand_sessionkey := space_id + strconv.Itoa(rand.Intn(2^10))
	valid := false
	for !valid { //혹시나 같은 것을 찾으면 다시 랜덤 돌리기위함
		for _, v := range *users {
			if rand_sessionkey == v.User_sessionkey {
				rand_sessionkey = Random_sessionkey_generator(space_id)
			}
		}
		valid = true
	}
	return rand_sessionkey
}
