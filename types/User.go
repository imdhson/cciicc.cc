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

func (users *Users) Remove_user(idx int) {
	//지워야될 곳(*users)[i]
	tmp := *users
	*users = Users{}
	for i := 0; i < len(tmp); i++ {
		if i == idx && i > 0 { //삭제할 것을 찾았을 떄
			i--
		} else if i <= 0 {

		} else {
			*users = append(*users, tmp[i])
		}
	}
}
