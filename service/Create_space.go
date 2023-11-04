package service

import (
	"ub/types"
)

func Create_space(space_name string) string {
	spaces := types.GetInstance_space()
	space_id := Random_space_id_generator()
	space := types.Space{Sp_id: space_id,
		Sp_name: space_name, Sp_view: 0}
	*spaces = append(*spaces, space)
	return space_id
}
