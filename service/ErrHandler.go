package service

import (
	"log"
)

func ErrHandler(e error, loc string) {
	if e != nil {
		log.Printf("%v=%v\n", loc, e)
	}
}
func CriticalErr(e error, loc string) {
	if e != nil {
		log.Panicf("%v=%v\n", loc, e)
	}
}
