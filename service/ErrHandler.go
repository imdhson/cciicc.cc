package service

import (
	"log"
)

func ErrHandler(e error) {
	log.Println("e: %v\n", e)
}
func CriticalErr(e error) {
	log.Printf("e: %v\n", e)
}
