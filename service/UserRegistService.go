package service

import "cciicc.cc/types"

func UserRegist(user types.User) bool {
	users := types.GetInstance_users()
	*users = append(*users, user)
	return true
}
