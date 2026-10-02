package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLogOpenFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test_log.txt")

	// Test creating the file
	file, err := LogOpenFile(filePath)
	if err != nil {
		t.Fatalf("LogOpenFile failed to create file: %v", err)
	}
	file.Close()

	// Test appending to the file
	file2, err := LogOpenFile(filePath)
	if err != nil {
		t.Fatalf("LogOpenFile failed to open existing file: %v", err)
	}
	defer file2.Close()

	// Verify it's writable
	_, err = file2.WriteString("test log")
	if err != nil {
		t.Errorf("Failed to write to log file: %v", err)
	}
}

func TestCreateFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test_create.txt")

	file, err := CreateFile(filePath)
	if err != nil {
		t.Fatalf("CreateFile failed: %v", err)
	}
	file.Close()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("File was not actually created")
	}

	// Test error case (invalid path)
	_, err = CreateFile(filepath.Join(tempDir, "nonexistent_dir", "test.txt"))
	if err == nil {
		t.Errorf("CreateFile should fail for invalid path")
	}
}

func TestOpenFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test_open.txt")

	// Create a file to open
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}
	f.Close()

	file, err := OpenFile(filePath)
	if err != nil {
		t.Fatalf("OpenFile failed: %v", err)
	}
	file.Close()

	// Test error case (file doesn't exist)
	_, err = OpenFile(filepath.Join(tempDir, "does_not_exist.txt"))
	if err == nil {
		t.Errorf("OpenFile should fail for non-existent file")
	}
}

// Helper to isolate tests with hardcoded paths
func setupWorkingDir(t *testing.T) {
	tempDir := t.TempDir()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	t.Cleanup(func() {
		os.Chdir(originalDir)
	})
}

func TestMakeDir_space_qr(t *testing.T) {
	setupWorkingDir(t)

	// Need to create parent directories first for os.Mkdir to work
	err := os.MkdirAll("wwwfiles/assets", 0755)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	err = MakeDir_space_qr()
	if err != nil {
		t.Fatalf("MakeDir_space_qr failed: %v", err)
	}

	info, err := os.Stat("wwwfiles/assets/space_qr")
	if os.IsNotExist(err) {
		t.Errorf("Directory was not created")
	} else if !info.IsDir() {
		t.Errorf("Expected a directory, got a file")
	}
}

func TestMakeDir_host_file(t *testing.T) {
	setupWorkingDir(t)

	// Need to create parent directories first for os.Mkdir to work
	err := os.MkdirAll("wwwfiles", 0755)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	err = MakeDir_host_file()
	if err != nil {
		t.Fatalf("MakeDir_host_file failed: %v", err)
	}

	info, err := os.Stat("wwwfiles/host_file")
	if os.IsNotExist(err) {
		t.Errorf("Directory was not created")
	} else if !info.IsDir() {
		t.Errorf("Expected a directory, got a file")
	}
}

func TestDeleteAll_space_qr(t *testing.T) {
	setupWorkingDir(t)

	// Setup: create directory and some files
	err := os.MkdirAll("wwwfiles/assets/space_qr", 0755)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}
	f, _ := os.Create("wwwfiles/assets/space_qr/test.png")
	f.Close()

	err = DeleteAll_space_qr()
	if err != nil {
		t.Fatalf("DeleteAll_space_qr failed: %v", err)
	}

	_, err = os.Stat("wwwfiles/assets/space_qr")
	if !os.IsNotExist(err) {
		t.Errorf("Directory should have been deleted")
	}
}

func TestDelete_space_qr(t *testing.T) {
	setupWorkingDir(t)

	// Setup
	err := os.MkdirAll("wwwfiles/assets/space_qr", 0755)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	spaceID := "test_space_123"
	filePath := filepath.Join("wwwfiles/assets/space_qr", spaceID+".png")

	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Setup failed to create file: %v", err)
	}
	f.Close()

	err = Delete_space_qr(spaceID)
	if err != nil {
		t.Fatalf("Delete_space_qr failed: %v", err)
	}

	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("File should have been deleted")
	}
}

func TestDelete_space_host_file(t *testing.T) {
	setupWorkingDir(t)

	// Setup
	err := os.MkdirAll("wwwfiles/host_file", 0755)
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	spaceID := "test_space_456"

	// Create a file that SHOULD be deleted
	shouldDelete := filepath.Join("wwwfiles/host_file", spaceID+".data.txt")
	f1, err := os.Create(shouldDelete)
	if err != nil {
		t.Fatalf("Setup failed to create file: %v", err)
	}
	f1.Close()

	// Create a file that SHOULD NOT be deleted
	shouldNotDelete := filepath.Join("wwwfiles/host_file", "other_space.data.txt")
	f2, err := os.Create(shouldNotDelete)
	if err != nil {
		t.Fatalf("Setup failed to create file: %v", err)
	}
	f2.Close()

	err = Delete_space_host_file(spaceID)
	if err != nil {
		t.Fatalf("Delete_space_host_file failed: %v", err)
	}

	// Verify target file is deleted
	_, err = os.Stat(shouldDelete)
	if !os.IsNotExist(err) {
		t.Errorf("Target file should have been deleted")
	}

	// Verify other file is NOT deleted
	_, err = os.Stat(shouldNotDelete)
	if os.IsNotExist(err) {
		t.Errorf("Other file should not have been deleted")
	}
}
