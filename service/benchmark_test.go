package service

import (
	"cciicc/types"
	"fmt"
	"testing"
)

func setupUsersLocal(numUsers int, spaceID string) *types.Users {
	users := &types.Users{}
	*users = make(types.Users, 0, numUsers)
	for i := 0; i < numUsers; i++ {
		*users = append(*users, types.User{
			User_name:            fmt.Sprintf("User%d", i),
			User_related_spaceid: spaceID,
		})
	}
	return users
}

func removeUsersRelatedSpaceIdValue(users *types.Users, space_id string) {
	n := 0
	for _, v := range *users {
		if v.User_related_spaceid != space_id {
			(*users)[n] = v
			n++
		}
	}
	for i := n; i < len(*users); i++ {
		(*users)[i] = types.User{} // Prevent memory leak
	}
	*users = (*users)[:n]
}

func removeUsersRelatedSpaceIdIndex(users *types.Users, space_id string) {
	usersSlice := *users
	n := 0
	for i := range usersSlice {
		if usersSlice[i].User_related_spaceid != space_id {
			usersSlice[n] = usersSlice[i]
			n++
		}
	}
	for i := n; i < len(usersSlice); i++ {
		usersSlice[i] = types.User{} // Prevent memory leak
	}
	*users = usersSlice[:n]
}

func BenchmarkRemoveUsersValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		users := setupUsersLocal(10000, "space1")
		b.StartTimer()
		removeUsersRelatedSpaceIdValue(users, "space1")
	}
}

func BenchmarkRemoveUsersIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		users := setupUsersLocal(10000, "space1")
		b.StartTimer()
		removeUsersRelatedSpaceIdIndex(users, "space1")
	}
}
