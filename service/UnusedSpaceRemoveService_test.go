package service

import (
	"cciicc/types"
	"fmt"
	"testing"
)

func setupUsers(numUsers int, spaceID string) {
	users := types.GetInstance_users()
	*users = make(types.Users, 0, numUsers)
	for i := 0; i < numUsers; i++ {
		*users = append(*users, types.User{
			User_name:            fmt.Sprintf("User%d", i),
			User_related_spaceid: spaceID,
		})
	}
}

func removeUsersRelatedSpaceIdOld(space_id string) {
	isRemove := false
	for !isRemove {
		users := types.GetInstance_users()
		for i, v := range *users {
			if v.User_related_spaceid == space_id {
				users.Remove_user(i)
				break
			}
		}
		isRemove = true
	}
}

func removeUsersRelatedSpaceIdNew(space_id string) {
	users := types.GetInstance_users()
	for i := len(*users) - 1; i >= 0; i-- {
		if (*users)[i].User_related_spaceid == space_id {
			users.Remove_user(i)
		}
	}
}

func BenchmarkRemoveUsersOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		setupUsers(1000, "space1")
		b.StartTimer()
		removeUsersRelatedSpaceIdOld("space1")
	}
}

func BenchmarkRemoveUsersNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		setupUsers(1000, "space1")
		b.StartTimer()
		removeUsersRelatedSpaceIdNew("space1")
	}
}
