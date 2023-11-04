package service

import "ub/types"

func UserRegist(user types.User) bool {
	users := types.GetInstance_users()
	*users = append(*users, user)
	return true
}
