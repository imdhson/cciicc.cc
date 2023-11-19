package types

import (
	"time"
)

type Spaces []Space

var spaces *Spaces

func GetInstance_spaces() *Spaces {
	if spaces == nil {
		spaces = &Spaces{}
	} else {
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
	Sp_c_id        int
	Sp_c_content   string
	Sp_c_guestname string
	Sp_c_color     Sp_c_color
	//Sp_c_comment   []Sp_comment //recursive 처리 - UI 복잡해서 일단 안하기로함
}
