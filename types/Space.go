package types

import (
	"fmt"
	"time"
)

type Spaces []Space

var spaces *Spaces

type Space struct {
	Sp_id           string
	Sp_name         string
	Sp_view         int
	Sp_lastupdate   time.Time
	Sp_chats        []Sp_chat
	Sp_file_status  Sp_file_status
	Sp_file_ext     string
	Sp_file_name    string
	Sp_file_context int // 파일의 위치: 초단위시각, 페이지 등
}

type Sp_chat struct {
	Sp_c_id int
	// Sp_c_rate      int //좋아요 싫어요 기능
	Sp_c_content   string
	Sp_c_guestname string
	// Sp_c_color     Sp_c_color
}

func GetInstance_spaces() *Spaces {
	if spaces == nil {
		spaces = &Spaces{}
	}
	return spaces
}

func (spaces *Spaces) Remove_space(idx int) {
	//지워야될 곳(*spaces)[i]
	old := *spaces
	*spaces = Spaces{}
	for i := 0; i < len(old); i++ {
		if i == idx { //삭제할 것을 찾았을 때
			fmt.Println("삭제중 space", i)
		} else {
			*spaces = append(*spaces, old[i])
		}
	}
}

// type Sp_c_color int

// const (
// 	WHILE  Sp_c_color = -1
// 	ORANGE Sp_c_color = iota
// 	SKYBLUE
// 	GREEN
// 	PINK
// 	RED
// 	BLUE
// )

type Sp_file_status int

const (
	SP_FILESTATUS_NONE Sp_file_status = iota
	SP_FILESTATUS_PDF
	SP_FILESTATUS_AUDIO
	SP_FILESTATUS_IMAGE
	SP_FILESTATUS_VIDEO
	SP_FILESTATUS_TEXT
)
