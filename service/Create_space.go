package service

import (
	"time"

	"cciicc.cc/types"
)

func Create_space(space_name string) string {
	spaces := types.GetInstance_spaces()
	space_id := Random_space_id_generator()
	space := types.Space{Sp_id: space_id,
		Sp_name: space_name, Sp_view: 0, Sp_lastupdate: time.Now()}
	*spaces = append(*spaces, space)
	return space_id
}
