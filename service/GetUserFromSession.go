package service

import (
	"cciicc/types"
)

func GetUserFromSession(session string) (types.User, bool) {
	users := types.GetInstance_users()
	// Dereference slice pointer once outside the loop
	usersSlice := *users
	// Optimization: Iterate backwards to find recently active users faster
	for i := len(usersSlice) - 1; i >= 0; i-- {
		if session == usersSlice[i].User_sessionkey {
			return usersSlice[i], true
		}
	}
	return types.User{}, false
}
