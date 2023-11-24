package types

import (
	"time"
)

type Spaces []Space

var spaces *Spaces

type Space struct {
	Sp_id         string
	Sp_name       string
	Sp_view       int
	Sp_lastupdate time.Time
	Sp_comments   []Sp_comment
}

type Sp_comment struct {
	Sp_c_id        int
	Sp_c_rate      int //좋아요 싫어요 기능
	Sp_c_content   string
	Sp_c_guestname string
	Sp_c_color     Sp_c_color
	//Sp_c_comment   []Sp_comment //recursive 처리 - UI 복잡해서 일단 안하기로함
}

func GetInstance_spaces() *Spaces {
	if spaces == nil {
		spaces = &Spaces{}
	} else {
	}
	return spaces
}

func (spaces *Spaces) Remove_space(idx int) {
	//지워야될 곳(*spaces)[i]
	tmp := *spaces
	*spaces = Spaces{}
	for i := 0; i < len(tmp)-1; i++ {
		if i == idx { //삭제할 것을 찾았을 떄
			i--
		} else {
			*spaces = append(*spaces, tmp[i])
		}
	}
}
