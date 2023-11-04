package main

import (
	"log"
	"net/http"
	"strconv"
	"ub/controller"
	"ub/service"
)

const PORT = 80
const SSLPORT = 443

func main() {
	//로그 기록 시작
	service.Logger()

	const PORT int = 80
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.URLHandler)
	log.Println(""+strconv.Itoa(PORT), "포트에서 요청을 기다리는 중...")
	go controller.CLI()
	//err := http.ListenAndServeTLS(":"+strconv.Itoa(SSLPORT), "ssl/combined.crt", "ssl/private.key", server)
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux) //암호화없음
	service.CriticalErr(err)
}
