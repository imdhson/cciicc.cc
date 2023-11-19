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

}
