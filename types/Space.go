package types

import (
	"log"
	"time"
)

type Spaces []Space

var spaces *Spaces

func GetInstance_space() *Spaces {
	if spaces == nil {
		log.Println("space 싱글톤 패턴 생성")
		spaces = &Spaces{}
	} else {
		log.Println("space 싱글톤 패턴 이미 생성되어 반환만 함.")
	}
	return spaces
}

type Space struct {
	Sp_id         string
	Sp_name       string
	Sp_view       int
	Sp_lastupdate time.Time
	Sp_comments   []Sp_comment
}

type Sp_comment struct {
	Sp_c_content   string
	Sp_c_guestname string
	Sp_c_color     int
	Sp_c_comment   []Sp_comment
}
