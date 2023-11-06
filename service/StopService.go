package service

import (
	"log"
	"os"
	"os/signal"
)

func StopService() {
	log.Println("서버를 중지하는 중...")
	//종료 전 파일 저장등 전처리 필요
	os.Exit(0)
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
