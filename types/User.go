package types

import "fmt"

type User struct {
	User_name            string
	User_sessionkey      string
	User_isHost          bool
	User_related_spaceid string
}
type Users []User

var users *Users

func GetInstance_users() *Users {
	if users == nil {
		users = &Users{}
	}
	return users
}

func (users *Users) Remove_user(idx int) {
	//지워야될 곳(*users)[i]
	old := *users
	*users = Users{}
	for i := 0; i < len(old); i++ {
		if i == idx { //삭제할 것을 찾았을 떄
			fmt.Println("삭제중 user", i)
		} else {
			*users = append(*users, old[i])
		}
	}
}
