package service

import (
	"cciicc.cc/types"
)

func GetUserFromSession(session string) (types.User, bool) {
	users := types.GetInstance_users()
	var rst types.User
	rst_success := false
	for _, v := range *users {
		if session == v.User_sessionkey {
			rst_success = true
			rst = v
		}
	}
	return rst, rst_success
}
