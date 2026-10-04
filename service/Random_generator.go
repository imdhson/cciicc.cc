package service

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"cciicc/types"
)

var cryptoRandIntn = func(max int64) int {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		// Fallback or panic; in standard environments, rand.Reader should not fail.
		panic(fmt.Sprintf("crypto/rand failed: %v", err))
	}
	return int(n.Int64())
}

func Random_space_id_generator() string {
	spaces := types.GetInstance_spaces()
	spacesSlice := *spaces
	valid := false
	var space_id string

	for !valid { //혹시나 같은 것을 찾으면 다시 랜덤 돌리기위함
		rand_int := cryptoRandIntn(999)
		rand_char := string(rune(cryptoRandIntn(26) + 97))
		space_id = rand_char + strconv.Itoa(rand_int)
		valid = true

		for i := len(spacesSlice) - 1; i >= 0; i-- { //순회하며 아이디같은지 찾고, 찾으면 랜덤
			if space_id == spacesSlice[i].Sp_id {
				valid = false
				break
			}
		}
	}
	return space_id
}

func Random_sessionkey_generator(space_id string) string {
	users := types.GetInstance_users()
	usersSlice := *users
	var rand_sessionkey string
	valid := false
	for !valid { //혹시나 같은 것을 찾으면 다시 랜덤 돌리기위함
		rand_sessionkey = space_id + strconv.Itoa(cryptoRandIntn(1<<16))
		valid = true
		for i := len(usersSlice) - 1; i >= 0; i-- {
			if rand_sessionkey == usersSlice[i].User_sessionkey {
				valid = false
				break
			}
		}
	}
	//해쉬함수
	hash := md5.New()
	hash.Write([]byte(rand_sessionkey))
	hashSum := hash.Sum(nil)
	sessionkeyStr := fmt.Sprintf("%x", hashSum)
	return sessionkeyStr
}
