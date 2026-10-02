package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteSpaceHostFile(t *testing.T) {
	// Setup temporary directory structure
	baseDir := "wwwfiles/host_file"
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create base directory: %v", err)
	}
	defer os.RemoveAll("wwwfiles") // Clean up after tests

	t.Run("Successfully deletes matching files", func(t *testing.T) {
		spaceID := "testspace"

		// Create a matching file
		matchingFile := filepath.Join(baseDir, spaceID+".txt")
		if err := os.WriteFile(matchingFile, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create matching file: %v", err)
		}

		// Create a non-matching file
		nonMatchingFile := filepath.Join(baseDir, "otherspace.txt")
		if err := os.WriteFile(nonMatchingFile, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create non-matching file: %v", err)
		}

		err := Delete_space_host_file(spaceID)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		// Verify matching file is deleted
		if _, err := os.Stat(matchingFile); !os.IsNotExist(err) {
			t.Errorf("Expected matching file to be deleted, but it still exists")
		}

		// Verify non-matching file still exists
		if _, err := os.Stat(nonMatchingFile); os.IsNotExist(err) {
			t.Errorf("Expected non-matching file to exist, but it was deleted")
		}
	})

	t.Run("Returns error when walk fails", func(t *testing.T) {
		// To simulate walk error, we can point to a non-existent directory
		// The original code hardcodes "wwwfiles/host_file/", so we have to manipulate that directory

		// Temporarily move the directory out of the way
		err := os.Rename("wwwfiles", "wwwfiles_temp")
		if err != nil {
			t.Fatalf("Failed to rename wwwfiles: %v", err)
		}

		err = Delete_space_host_file("testspace")
		if err == nil {
			t.Errorf("Expected an error when directory doesn't exist")
		}

		// Restore the directory
		err = os.Rename("wwwfiles_temp", "wwwfiles")
		if err != nil {
			t.Fatalf("Failed to restore wwwfiles: %v", err)
		}
	})

	t.Run("Returns error when file deletion fails", func(t *testing.T) {
		spaceID := "testspace_error"

		// Create a matching file
		matchingFile := filepath.Join(baseDir, spaceID+".txt")
		if err := os.WriteFile(matchingFile, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create matching file: %v", err)
		}

		// Make the directory read-only to prevent deletion
		err = os.Chmod(baseDir, 0555)
		if err != nil {
			t.Fatalf("Failed to change directory permissions: %v", err)
		}

		err = Delete_space_host_file(spaceID)
		if err == nil {
			t.Errorf("Expected an error when file deletion is prevented")
		}

		// Restore permissions
		err = os.Chmod(baseDir, 0755)
		if err != nil {
			t.Fatalf("Failed to restore directory permissions: %v", err)
		}
	})
}
