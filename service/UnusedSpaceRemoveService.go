package service

import (
	"time"

	"cciicc.cc/storage"
	"cciicc.cc/types"
)

func UnusedSpaceRemoveService() {
	for {
		spaces := types.GetInstance_spaces()
		isRemove := false
		for i, v := range *spaces {
			lastupdateSubtract_hour := v.Sp_lastupdate.Add(time.Hour * -1)
			timeCompare := time.Now().Compare(lastupdateSubtract_hour) //-1: 1시간 이내 1: 1시간 지나서 삭제
			if timeCompare == 1 {
				spaces.Remove_space(i) //space [i] 삭제

				file_remove_err := storage.Delete_space_ar("wwwfiles/assets/space_qr/" + v.Sp_id + ".png")
				CriticalErr(file_remove_err, "UnusedSpaceRemoveService - 파일 삭제")

				isRemove = true
				break
			}
		}
		if !isRemove { //리무브를 했으면 break하기때문에 sleep를 안하고 즉시 반복수행
			time.Sleep(time.Hour * 1) //1시간 마다 순회하며 수행
			isRemove = false
		}
	}

}
