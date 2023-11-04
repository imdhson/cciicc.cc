package types

import "log"

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
		log.Println("user 싱글톤 패턴 생성")
		users = &Users{}
	} else {
		log.Println("user 싱글톤 패턴 이미 생성됨.")
	}
	return users
}
