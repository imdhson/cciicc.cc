package controller

import (
	"fmt"
	"ub/service"
	"ub/types"
)

func CLI() {
	var t string
	for t != "stop" {
		fmt.Print("UFB > ")
		fmt.Scanln(&t)
		switch t {
		case "stop", "exit", "quit":
			service.StopService()
		case "data", "test", "datas":
			fmt.Println("\033[2J")
			spaces := types.GetInstance_spaces()
			fmt.Println("------spaces------")
			for _, v := range *spaces { //스페이스공간 디버그용
				fmt.Println(v.Sp_id + "|" + v.Sp_name)
			}
			fmt.Println("------users------")
			users := types.GetInstance_users()
			for _, v := range *users { //스페이스공간 디버그용
				fmt.Println(v.User_name + "|" + v.User_sessionkey)
			}
			fmt.Println()
		case "clear":
			fmt.Println("\033[2J")
		case "cleardata":
			service.ClearDatas()
		case "spaces":
			spaces := types.GetInstance_spaces()
			for _, v := range *spaces { //스페이스공간 디버그용
				fmt.Println(v.Sp_id + "|" + v.Sp_name)
			}
		case "users":
			users := types.GetInstance_users()
			for _, v := range *users { //스페이스공간 디버그용
				fmt.Println(v.User_name + "|" + v.User_sessionkey)
			}
		default:
			fmt.Println("invalid 명령어")
		}

	}
}
