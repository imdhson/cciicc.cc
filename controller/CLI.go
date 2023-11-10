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
		case "spaces_all":
			spaces := types.GetInstance_spaces()
			fmt.Println(spaces)
		case "users_all":
			users := types.GetInstance_users()
			fmt.Println(users)
		case "users":
			users := types.GetInstance_users()
			for _, v := range *users { //스페이스공간 디버그용
				fmt.Println(v.User_name + "|" + v.User_sessionkey)
			}
		case "save", "savedata":
			service.SaveDatas()
		case "1":
			comment := types.Sp_comment{
				Sp_c_content:   "form_comment",
				Sp_c_guestname: "user.User_name",
				Sp_c_color:     1000,
			}
			//spaces := types.GetInstance_spaces()
			//(*spaces)[0].Sp_comments = append((*spaces)[0].Sp_comments, comment)
			space, _ := service.GetSpaceFrom_space_id("s835")
			space.Sp_comments = append(space.Sp_comments, comment)
		default:
			fmt.Println("invalid 명령어")
		}
	}
}
