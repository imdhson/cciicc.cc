package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"ub/controller"
	"ub/service"
)

func main() {
	// 로그 기록
	// 로그 파일 생성
	log_f, log_err := os.OpenFile("last.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	service.CriticalErr(log_err)
	defer log_f.Close()

	// 로그 출력 설정
	log.SetOutput(io.MultiWriter(os.Stdout, log_f))

	const PORT int = 80
	server := http.NewServeMux()
	server.Handle("/", http.HandlerFunc(controller.IndexHandler))
	server.Handle("/host", http.HandlerFunc(controller.HostHandler))
	log.Println(":"+strconv.Itoa(PORT), "에서 요청을 기다리는 중:")
	//err := http.ListenAndServeTLS(":"+strconv.Itoa(PORT), "sslforfree/combined.crt", "sslforfree/private.key", server)
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), server)
	service.CriticalErr(err)
}
