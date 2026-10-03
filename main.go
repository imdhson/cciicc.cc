package main

import (
	"log"
	"net/http"
	"os"

	"cciicc/controller"
	"cciicc/service"

	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {

	service.StartService()
	go service.DetectStopService() //Ctrl+C (인터럽트)시 종료 서비스 호출
	// go controller.CLI() //Command line 인터페이스 호출 : 필요시 주석 해제

	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.URLHandler)

	port := getEnv("PORT", "80")
	sslPort := getEnv("SSL_PORT", "443")
	enableTLS := getEnv("ENABLE_TLS", "false")
	tlsCert := getEnv("TLS_CERT", "certkey")
	tlsKey := getEnv("TLS_KEY", "key")

	var csrfAuthKey []byte
	if key := os.Getenv("CSRF_AUTH_KEY"); len(key) == 32 {
		csrfAuthKey = []byte(key)
	} else {
		csrfAuthKey = securecookie.GenerateRandomKey(32)
	}

	isSecure := false
	if enableTLS == "true" {
		isSecure = true
	}

	csrfOpts := []csrf.Option{
		csrf.Secure(isSecure),
	}

	if trustedOrigin := os.Getenv("CSRF_TRUSTED_ORIGIN"); trustedOrigin != "" {
		csrfOpts = append(csrfOpts, csrf.TrustedOrigins([]string{trustedOrigin}))
	}

	CSRF := csrf.Protect(csrfAuthKey, csrfOpts...)
	handler := CSRF(mux)

	if enableTLS == "true" {
		go func() {
			log.Println("TLS 서버 시작: 포트 " + sslPort + "에서 요청을 기다리는 중...")
			err := http.ListenAndServeTLS(":"+sslPort, tlsCert, tlsKey, handler)
			if err != nil {
				log.Printf("TLS 서버 에러: %v\n", err)
			}
		}()
	}

	log.Println("HTTP 서버 시작: 포트 " + port + "에서 요청을 기다리는 중...")
	err := http.ListenAndServe(":"+port, handler) //암호화없음
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
