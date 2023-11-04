package types

import (
	"log"
)

// var spaces *Spaces
// var once sync.Once

// 	func GetInstance() *Spaces {
// 		if spaces == nil {
// 			once.Do(func() {
// 				log.Println("싱글톤 패턴 생성")
// 				spaces = &Spaces{}
// 			})
// 		} else {
// 			log.Println("싱글톤 이미 생성됨")
// 		}
// 		return spaces
// 	}

var spaces *Spaces

func GetInstance() *Spaces {
	if spaces == nil {
		log.Println("싱글톤 패턴 생성")
		spaces = &Spaces{}
	} else {
		log.Println("싱글톤 패턴 이미 생성되어 반환만 함.")
	}
	return spaces
}
