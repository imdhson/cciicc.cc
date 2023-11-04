package controller

import (
	"fmt"
	"ub/service"
)

func CLI() {
	var t string
	for t != "stop" {
		fmt.Print("UFB > ")
		fmt.Scanln(&t)
		switch t {
		case "stop", "exit", "quit":
			service.StopService()
		}
	}
}
