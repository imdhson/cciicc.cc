package service

import (
	"io"
	"log"
	"os"
	"ub/storage"
)

func Logging(i string) {
	//ip 넣기
	log.Println(i)
}

func Logger() {
	// 로그 기록
	// 로그 파일 생성
	log_f, log_err := storage.LogOpenFile("server.log")
	ErrHandler(log_err, "log_err")

	// 로그 출력 설정
	log.SetOutput(io.MultiWriter(os.Stdout, log_f))
}
