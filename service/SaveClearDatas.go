package service

import (
	"encoding/json"
	"fmt"
	"ub/storage"
	"ub/types"
)

func SaveDatas() {
	fmt.Println("데이터 저장 시작")
	// users 저장
	users := types.GetInstance_users()
	file_users, f_users_err := storage.CreateFile("ub_users.data")
	ErrHandler(f_users_err, "f_users_err createFile")
	defer file_users.Close()

	user_encoded, err := json.MarshalIndent(users, " ", "	")
	ErrHandler(err, "user_encoded")
	file_users.Write(user_encoded)

	spaces := types.GetInstance_spaces()
	file_spaces, f_spaces_err := storage.CreateFile("ub_spaces.data")
	ErrHandler(f_spaces_err, "f_spaces_err CreateFile")
	defer file_spaces.Close()

	spaces_encoded, err := json.MarshalIndent(spaces, " ", "	")
	ErrHandler(err, "spaces_encoded")
	file_spaces.Write(spaces_encoded)
	fmt.Println("데이터 저장 완료")
}

func ClearDatas() {
	spaces := types.GetInstance_spaces()
	users := types.GetInstance_users()
	*spaces = nil
	*users = nil
	fmt.Println(users)
}
