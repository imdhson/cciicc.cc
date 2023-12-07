package service

import (
	"cciicc/types"
)

func GetUserFromSession(session string) (types.User, bool) {
	users := types.GetInstance_users()
	var rst types.User
	for _, v := range *users {
		if session == v.User_sessionkey {
			rst = v
			return rst, true
		}
	}
	return rst, false
}
