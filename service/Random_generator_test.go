package service

import (
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
