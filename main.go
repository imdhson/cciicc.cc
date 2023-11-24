package main

import (
	"log"
	"net/http"
	"strconv"

	"cciicc.cc/controller"
	"cciicc.cc/service"
)

const PORT = 80
const SSLPORT = 443

func main() {

	//이전 데이터 불러오기
	service.StartService()
	go service.DetectStopService() //Ctrl+C (인터럽트)시 종료 서비스 호출
	go controller.CLI()            //Command line 인터페이스 호출

	const PORT int = 80
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.URLHandler)
	log.Println(""+strconv.Itoa(PORT), "포트에서 요청을 기다리는 중...")
	//err := http.ListenAndServeTLS(":"+strconv.Itoa(SSLPORT), "ssl/combined.crt", "ssl/private.key", mux)
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux) //암호화없음
	service.CriticalErr(err, "http.ListenAndServe")

	//next up:
	// 자동 데이터 저장 []

	// 글 많아지면 맨 아래로 자동스크롤 only sort by id [완료]
	// 인터랙션 하루동안 없으면 스페이스, qr 자동 삭제 - for stability [완료 테스트중]
	// 삭제되면 네트워크 유실 에러 발생[ 완료 테스트중 ]

	// 좋아요순 구현 시 인터랙션 안하면 새로고침하며 새로 그리기 [완료 테스트중]
	// form으로 보내기를 js로 받아와서 비동기식으로 처리 [완료 테스트중]
	// 좋아요 싫어요 js 로 전송  [완료]
	// 정렬 옵션 id순 or 좋아요 순 [완료]
	// 댓글 색상 처리 [완료]
}
