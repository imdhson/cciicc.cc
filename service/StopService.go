package service

import "os"

func StopService() {
	//종료 전 파일 저장등 전처리 필요
	os.Exit(0)
}
