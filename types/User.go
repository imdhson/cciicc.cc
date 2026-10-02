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
	if idx < 0 || idx >= len(*users) {
		return
	}
	fmt.Println("삭제중 user", idx)
	*users = append((*users)[:idx], (*users)[idx+1:]...)
}
