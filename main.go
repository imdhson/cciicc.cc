package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"cciicc/controller"
	"cciicc/service"
	"cciicc/types"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default values")
	} else {
		if url := os.Getenv("URL_ADDESS"); url != "" {
			types.URL_ADDESS = url
		}
	}

	// Load messages.json
	messagesFile, err := os.Open("messages.json")
	if err == nil {
		defer messagesFile.Close()
		var msgs map[string]string
		if err := json.NewDecoder(messagesFile).Decode(&msgs); err == nil {
			if v, ok := msgs["SERVICE_NAME"]; ok { types.SERVICE_NAME = v }
			if v, ok := msgs["SERVICE_DETAIL"]; ok { types.SERVICE_DETAIL = v }
			if v, ok := msgs["FOOTER_TERMS"]; ok { types.FOOTER_TERMS = v }
			if v, ok := msgs["MAIN_NEXT_UP"]; ok { types.MAIN_NEXT_UP = v }
			if v, ok := msgs["MAIN_HOST"]; ok { types.MAIN_HOST = v }
			if v, ok := msgs["MAIN_GUEST"]; ok { types.MAIN_GUEST = v }
			if v, ok := msgs["HOST_DETAIL"]; ok { types.HOST_DETAIL = v }
			if v, ok := msgs["HOST_SPACENAME"]; ok { types.HOST_SPACENAME = v }
			if v, ok := msgs["HOST_SPACENAME_INPUT"]; ok { types.HOST_SPACENAME_INPUT = v }
			if v, ok := msgs["HOST_USERNAME"]; ok { types.HOST_USERNAME = v }
			if v, ok := msgs["HOST_USERNAME_INPUT"]; ok { types.HOST_USERNAME_INPUT = v }
			if v, ok := msgs["HOST_FORM_BUTTON"]; ok { types.HOST_FORM_BUTTON = v }
			if v, ok := msgs["GUEST_DETAIL"]; ok { types.GUEST_DETAIL = v }
			if v, ok := msgs["GUEST_SPACEID"]; ok { types.GUEST_SPACEID = v }
			if v, ok := msgs["GUEST_SPACEID_INPUT"]; ok { types.GUEST_SPACEID_INPUT = v }
			if v, ok := msgs["GUEST_USERNAME"]; ok { types.GUEST_USERNAME = v }
			if v, ok := msgs["GUEST_USERNAME_INPUT"]; ok { types.GUEST_USERNAME_INPUT = v }
			if v, ok := msgs["GUEST_FORM_BUTTON"]; ok { types.GUEST_FORM_BUTTON = v }
			if v, ok := msgs["CONTENT_VIEW_COUNT"]; ok { types.CONTENT_VIEW_COUNT = v }
			if v, ok := msgs["CONTENT_ORDER"]; ok { types.CONTENT_ORDER = v }
			if v, ok := msgs["CONTENT_SEND"]; ok { types.CONTENT_SEND = v }
			if v, ok := msgs["ERROR_TITLE"]; ok { types.ERROR_TITLE = v }
			if v, ok := msgs["ERROR_CONTENT"]; ok { types.ERROR_CONTENT = v }
			if v, ok := msgs["ERROR_MAIN"]; ok { types.ERROR_MAIN = v }
		} else {
			log.Println("Error parsing messages.json")
		}
	} else {
		log.Println("messages.json not found, using default values")
	}
}

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
