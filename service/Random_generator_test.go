package service

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"testing"

	"cciicc/types"
)

func TestRandomSpaceIdGeneratorFormat(t *testing.T) {
	spaces := types.GetInstance_spaces()
	// Clear any existing spaces to avoid side effects
	*spaces = types.Spaces{}

	pattern := `^[a-z][0-9]{1,3}$`
	re := regexp.MustCompile(pattern)

	for i := 0; i < 100; i++ {
		id := Random_space_id_generator()
		if !re.MatchString(id) {
			t.Errorf("Generated ID %s does not match expected format %s", id, pattern)
		}
	}
}

func TestRandomSpaceIdGeneratorCollision(t *testing.T) {
	spaces := types.GetInstance_spaces()

	// Clear any existing spaces
	*spaces = types.Spaces{}

	// Generate some IDs and add them to spaces
	existingIDs := make(map[string]bool)
	for i := 0; i < 100; i++ {
		id := Random_space_id_generator()
		existingIDs[id] = true
		*spaces = append(*spaces, types.Space{Sp_id: id})
	}

	// Generate a new ID and verify it does not collide with any existing ones
	for i := 0; i < 100; i++ {
		newID := Random_space_id_generator()
		if existingIDs[newID] {
			t.Errorf("Generated ID %s collided with existing ID in spaces", newID)
		}
		// Add to spaces to continue testing collision
		existingIDs[newID] = true
		*spaces = append(*spaces, types.Space{Sp_id: newID})
	}

	// Clean up
	*spaces = types.Spaces{}
}

func TestRandomSessionkeyGeneratorFormat(t *testing.T) {
	users := types.GetInstance_users()
	*users = types.Users{}

	space_id := "testspace"
	key := Random_sessionkey_generator(space_id)

	// MD5 hash in hex is exactly 32 characters long
	if len(key) != 32 {
		t.Errorf("Expected session key length 32, got %d", len(key))
	}

	matched, _ := regexp.MatchString("^[0-9a-f]{32}$", key)
	if !matched {
		t.Errorf("Generated session key %s is not a valid MD5 hex string", key)
	}
}

func TestRandomSessionkeyGeneratorCollision(t *testing.T) {
	// Setup custom rand function to force collision
	originalCryptoRandIntn := cryptoRandIntn
	defer func() {
		cryptoRandIntn = originalCryptoRandIntn
	}()

	users := types.GetInstance_users()
	*users = types.Users{}

	space_id := "testspace"

	// Mock sequence
	seq := []int{100, 100, 200}
	idx := 0
	cryptoRandIntn = func(max int64) int {
		if idx < len(seq) {
			val := seq[idx]
			idx++
			return val
		}
		return 300 // fallback
	}

	// Set up collision condition
	// 1st generated unhashed key: "testspace" + "100" = "testspace100"
	*users = append(*users, types.User{User_sessionkey: "testspace100"})

	// Expect 2nd generation: "testspace100" -> collides again
	// Expect 3rd generation: "testspace200" -> valid
	// Should break loop and md5 hash "testspace200"

	key := Random_sessionkey_generator(space_id)

	// verify that idx advanced to 3 (meaning it called cryptoRandIntn 3 times: initial + 2 collisions checked)
	if idx != 3 {
		t.Errorf("Expected cryptoRandIntn to be called 3 times, got %d", idx)
	}

	// Verify the final key generated is the MD5 of "testspace200"
	hash := md5.New()
	hash.Write([]byte("testspace200"))
	expectedKey := fmt.Sprintf("%x", hash.Sum(nil))

	if key != expectedKey {
		t.Errorf("Expected session key %s (MD5 of testspace200), got %s", expectedKey, key)
	}
}
