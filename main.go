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
	// go controller.CLI() //Command line 인터페이스 호출 : 필요시 주석 해제

	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.URLHandler)
	go http.ListenAndServeTLS(":"+strconv.Itoa(SSLPORT), "certkey", "key", mux)
	log.Println(""+strconv.Itoa(SSLPORT), "포트에서 요청을 기다리는 중...")

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux) //암호화없음
	log.Println(""+strconv.Itoa(PORT), "포트에서 요청을 기다리는 중...")
	service.CriticalErr(err, "http.ListenAndServe")

	//next up:

	// space 삭제시 파일 삭제, ws_space 삭제
	// 새로 추가한 함수들의 모듈화(컨트롤러에서 서비스로 이동)
	// 채팅 기능 전체 열람 기능 추가
	// context 수신 켜기 끄기 구현
	// 인터랙션 하루동안 없으면 스페이스, qr 자동 삭제 - for stability 오류발견 0번쩨 인덱스 삭제가 안됨 [완료, 테스트중]
	// 다수에 사용자가 동시 사용시 계정 섞이는 문제 [테스트 중]

	// 채팅 내용 표시 기능 추가 [완료]
	//space 자동 삭제시 관련된 user들 자동 삭제[완료]
	// 자동 데이터 저장 [완료]
	// 글 많아지면 맨 아래로 자동스크롤 only sort by id [완료]
}
