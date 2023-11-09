package types

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
	} else {
	}
	return users
}
