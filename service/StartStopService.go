package service

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/signal"
	"ub/types"
)

func StopService() {
	log.Println("서버를 중지하는 중...")

	//users 저장
	users := types.GetInstance_users()
	file_users, f_users_err := os.Create("ub_users.data")
	ErrHandler(f_users_err)
	defer file_users.Close()
	user_encoded, err := json.MarshalIndent(users, " ", "	")
	ErrHandler(err)
	file_users.Write(user_encoded)

	spaces := types.GetInstance_spaces()
	file_spaces, f_spaces_err := os.Create("ub_spaces.data")
	ErrHandler(f_spaces_err)
	defer file_spaces.Close()
	spaces_encoded, err := json.MarshalIndent(spaces, " ", "	")
	ErrHandler(err)
	file_spaces.Write(spaces_encoded)

	os.Exit(0)
}

func StartService() {
	log.Println("이전에 저장된 데이터를 불러오는 중...")
	//users 불러오기
	file_users, f_users_err := os.Open("ub_users.data")
	ErrHandler(f_users_err)
	defer file_users.Close()
	user_data, err := io.ReadAll(file_users)
	ErrHandler(err)
	users := types.GetInstance_users()
	err = json.Unmarshal(user_data, &users)
	ErrHandler(err)
	if err == nil {
		log.Println("users 저장된 데이터 불러오기 성공.")
	} else {
		log.Println("users 저장된 데이터 불러오기 실패")
	}

	//spaces 불러오기
	file_spaces, f_spaces_err := os.Open("ub_spaces.data")
	ErrHandler(f_spaces_err)
	defer file_users.Close()
	spaces_data, err := io.ReadAll(file_spaces)
	ErrHandler(err)
	spaces := types.GetInstance_spaces()
	err = json.Unmarshal(spaces_data, &spaces)
	ErrHandler(err)
	if err == nil {
		log.Println("spaces 저장된 데이터 불러오기 성공.")
	} else {
		log.Println("spaces 저장된 데이터 불러오기 실패")
	}
}

func DetectStopService() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		select {
		case <-c:
			log.Println()
			StopService()
		default:
		}
	}
}
