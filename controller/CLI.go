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
		case "test":
			spaces := types.GetInstance_space()
			for _, v := range *spaces { //스페이스공간 디버그용
				fmt.Println(v.Sp_id + "|" + v.Sp_name)
			}
			users := types.GetInstance_users()
			for _, v := range *users { //스페이스공간 디버그용
				fmt.Println(v.User_name + "|" + v.User_sessionkey)
			}
		case "clear":
			fmt.Println("\033[2J")
		}
	}
}
