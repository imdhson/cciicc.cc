package service

import (
	"log"
)

func ErrHandler(e error) {
	log.Println("e: %v\n", e)
}
func CriticalErr(e error) {
	log.Panic("e: %v\n", e)
}
