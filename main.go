package main

import (
	"log"
	"net/http"
	"strconv"

	"cciicc/controller"
	"cciicc/service"
)

const PORT = 80
const SSLPORT = 443

// you need to change the URL_ADDESS which is located in types.CONST.go

func main() {

	service.StartService()
	go service.DetectStopService() //Ctrl+C (인터럽트)시 종료 서비스 호출

	// go controller.CLI()            //Command line 인터페이스 호출 : 필요시 주석 해제

	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.URLHandler)
	go http.ListenAndServeTLS(":"+strconv.Itoa(SSLPORT), "certkey", "key", mux)
	log.Println(""+strconv.Itoa(SSLPORT), "포트에서 요청을 기다리는 중...")

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux) //암호화없음
	log.Println(""+strconv.Itoa(PORT), "포트에서 요청을 기다리는 중...")
	service.CriticalErr(err, "http.ListenAndServe")

	//next up:

	//space 자동 삭제시 관련된 user들 자동 삭제
	// 인터랙션 하루동안 없으면 스페이스, qr 자동 삭제 - for stability [완료 테스트중]

	// 자동 데이터 저장 [완료]
	// form으로 보내기를 js로 받아와서 비동기식으로 처리 [완료]
	// 좋아요순 구현 시 인터랙션 안하면 새로고침하며 새로 그리기 [완료]
	// 글 많아지면 맨 아래로 자동스크롤 only sort by id [완료]
	// 좋아요 싫어요 js 로 전송  [완료]
	// 정렬 옵션 id순 or 좋아요 순 [완료]
	// 댓글 색상 처리 [완료]
}
