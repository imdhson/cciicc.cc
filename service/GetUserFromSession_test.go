package service

import (
	"cciicc/types"
	"testing"
)

func TestGetUserFromSession(t *testing.T) {
	// Setup dummy user in the singleton
	users := types.GetInstance_users()
	// Clear the singleton to ensure a clean state
	*users = types.Users{}

	validSession := "valid_session_key_123"
	invalidSession := "invalid_session_key_456"

	dummyUser := types.User{
		User_name:            "Test User",
		User_sessionkey:      validSession,
		User_isHost:          false,
		User_related_spaceid: "space_123",
	}

	*users = append(*users, dummyUser)

	t.Run("Hit Case - Valid Session", func(t *testing.T) {
		user, found := GetUserFromSession(validSession)
		if !found {
			t.Errorf("Expected to find user with session %s, but did not", validSession)
		}
		if user.User_sessionkey != validSession {
			t.Errorf("Expected user session key %s, got %s", validSession, user.User_sessionkey)
		}
		if user.User_name != dummyUser.User_name {
			t.Errorf("Expected user name %s, got %s", dummyUser.User_name, user.User_name)
		}
	})

	t.Run("Miss Case - Invalid Session", func(t *testing.T) {
		_, found := GetUserFromSession(invalidSession)
		if found {
			t.Errorf("Expected to not find user with session %s, but did", invalidSession)
		}
	})
}
