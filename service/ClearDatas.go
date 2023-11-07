package service

import (
	"fmt"
	"ub/types"
)

func ClearDatas() {
	spaces := types.GetInstance_spaces()
	users := types.GetInstance_users()
	*spaces = nil
	*users = nil
	fmt.Println(users)
}
