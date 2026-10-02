package service

import (
	"cciicc/types"
)

func GetUserFromSession(session string) (types.User, bool) {
	users := types.GetInstance_users()
	for i := range *users {
		if session == (*users)[i].User_sessionkey {
			return (*users)[i], true
		}
	}
	return types.User{}, false
}
