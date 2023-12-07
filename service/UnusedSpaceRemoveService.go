package service

import (
	"log"
	"time"

	"cciicc/storage"
	"cciicc/types"
)

func UnusedSpaceRemoveService() {
	for { //종료시까지 무한 루프 돌아야해서 이 안에 코드 작성
		spaces := types.GetInstance_spaces()
		isRemove := false

		for i, v := range *spaces {
			lastupdate_add_hour := v.Sp_lastupdate.Add(time.Hour * types.AUTOREMOVE_UNUSED_SPACE_HOURS)

			timeCompare := time.Now().Compare(lastupdate_add_hour) //compare 반환 값 :: -1: 1시간 이내 1: 1시간 지나서 삭제

			if timeCompare == 1 {
				rm_space_id := (*spaces)[i].Sp_id
				log.Println("사용되지 않는 space 삭제: ", rm_space_id)

				file_remove_err := storage.Delete_space_ar("wwwfiles/assets/space_qr/" + v.Sp_id + ".png")
				CriticalErr(file_remove_err, "UnusedSpaceRemoveService - 파일 삭제")

				Remove_users_related_space_id(rm_space_id) //space_id와 연관된 유저들 삭제

				spaces.Remove_space(i) //space [i] 삭제

				isRemove = true
				break
			}
		}
		if !isRemove { //리무브를 했으면 break하기때문에 sleep를 안하고 즉시 반복수행
			log.Println("wait next for AUTOREMOVE_UNUSED_SPACE_HOURS ")
			time.Sleep(time.Hour * types.AUTOREMOVE_UNUSED_SPACE_HOURS) //1시간 마다 순회하며 수행
		}
	}
}

func Remove_users_related_space_id(space_id string) {

	users := types.GetInstance_users()
	isRemove := false
	for !isRemove {
		for i, v := range *users {
			if v.User_related_spaceid == space_id {
				log.Println("사용되지 않는 space의 user 삭제중:", v.User_name)
				users.Remove_user(i)
				break
			}
			isRemove = true
		}
	}

}
