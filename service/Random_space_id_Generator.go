package service

import (
	"math/rand"
	"strconv"
)

func Random_space_id_generator() string {
	rand_int := rand.Intn(999)
	rand_char := string(rune(rand.Intn(26) + 97))
	space_id := rand_char + strconv.Itoa(rand_int)
	return space_id
}
