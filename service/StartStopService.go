package service

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
	"ub/storage"
	"ub/types"
)

func StopService() {
	log.Println("서버를 중지하는 중...")
	SaveDatas()
	os.Exit(0)
}

func StartService() {
	// 로그 기록 시작
	Logger()

	log.Println("이전에 저장된 데이터를 불러오는 중...")
	//users 불러오기
	file_users, f_users_err := storage.OpenFile("ub_users.data")
	ErrHandler(f_users_err, "f_users_err OpenFile")
	defer file_users.Close()
	user_data, err := io.ReadAll(file_users)
	ErrHandler(err, "io.ReadAll(file_users)")
	users := types.GetInstance_users()
	err = json.Unmarshal(user_data, &users)
	ErrHandler(err, "json.Unmarshal(user_data, &users)")
	if err == nil {
		log.Println("users 저장된 데이터 불러오기 성공.")
	} else {
		log.Println("users 저장된 데이터 불러오기 실패")
	}

	//spaces 불러오기
	file_spaces, f_spaces_err := storage.OpenFile("ub_spaces.data")
	ErrHandler(f_spaces_err, "f_spaces_err OpenFile")
	defer file_users.Close()
	spaces_data, err := io.ReadAll(file_spaces)
	ErrHandler(err, "spaces_data, err := io.ReadAll(file_spaces)")
	spaces := types.GetInstance_spaces()
	err = json.Unmarshal(spaces_data, &spaces)
	ErrHandler(err, "json.Unmarshal(spaces_data, &spaces)")
	if err == nil {
		log.Println("spaces 저장된 데이터 불러오기 성공.")
	} else {
		log.Println("spaces 저장된 데이터 불러오기 실패")
	}

	//space_qr 정상 작동을 위해 디렉터리 생성 시도
	storage.MakeDir_space_qr()
}

func DetectStopService() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		time.Sleep(time.Second * 2)
		select {
		case <-c:
			log.Println()
			StopService()
		default:
		}
	}
}
