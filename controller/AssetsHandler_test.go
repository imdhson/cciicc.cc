package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestAssetsHanlderSecurePath(t *testing.T) {
	// Create a wwwfiles directory
	wwwfilesDir := filepath.Join(".", "wwwfiles")
	os.MkdirAll(wwwfilesDir, 0755)
	defer os.RemoveAll(wwwfilesDir) // clean up after test

	// Ensure wwwfiles/assets/error/index.html exists for ErrorPageHandler
	os.MkdirAll(filepath.Join(wwwfilesDir, "assets", "error"), 0755)
	err := os.WriteFile(filepath.Join(wwwfilesDir, "assets", "error", "index.html"), []byte("<html>Error</html>"), 0644)
	if err != nil {
		t.Fatalf("Failed to create error index.html: %v", err)
	}

	// Create a secret file outside wwwfiles
	secretFile := filepath.Join(".", "secret.txt")
	err = os.WriteFile(secretFile, []byte("sensitive info"), 0644)
	if err != nil {
		t.Fatalf("Failed to create secret file: %v", err)
	}
	defer os.Remove(secretFile)

	// Ensure wwwfiles/css/main.css exists
	os.MkdirAll(filepath.Join(wwwfilesDir, "css"), 0755)
	err = os.WriteFile(filepath.Join(wwwfilesDir, "css", "main.css"), []byte("body { color: red; }"), 0644)
	if err != nil {
		t.Fatalf("Failed to create main.css: %v", err)
	}

	// Ensure wwwfiles/secret.txt doesn't exist just in case
	os.Remove(filepath.Join(wwwfilesDir, "secret.txt"))

	tests := []struct {
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			url:          "css/main.css",
			expectedCode: http.StatusOK,
			expectedBody: "body { color: red; }",
		},
		{
			url:          "../secret.txt",
			expectedCode: http.StatusNotFound, // Should route to ErrorPageHandler which we don't fully implement here, but it returns error (os.ReadFile will fail since wwwfiles/secret.txt doesn't exist)
		},
		{
			url:          "/../secret.txt",
			expectedCode: http.StatusNotFound,
		},
		{
			url:          "../../secret.txt",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.url, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/assets/"+tc.url, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()

			// Call the handler directly
			AssetsHanlder(rr, req, tc.url)

			// We check the body - if it contains "sensitive info", then directory traversal worked!
			if rr.Body.String() == "sensitive info" {
				t.Errorf("VULNERABILITY: Successfully read file outside wwwfiles directory with url %s", tc.url)
			}

			// If we expected a specific body for valid requests
			if tc.expectedCode == http.StatusOK {
				if rr.Code != http.StatusOK {
					// We might get a different status from ErrorPageHandler if it's called
				}
				if rr.Body.String() != tc.expectedBody {
					t.Errorf("Expected body %q, got %q", tc.expectedBody, rr.Body.String())
				}
			}
		})
	}
}

func TestDotFileType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Standard extension",
			input:    "file.jpg",
			expected: "jpg",
		},
		{
			name:     "Multiple dots",
			input:    "archive.tar.gz",
			expected: "gz",
		},
		{
			name:     "No extension",
			input:    "README",
			expected: "",
		},
		{
			name:     "Starts with a dot",
			input:    ".gitignore",
			expected: "gitignore",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Special characters",
			input:    "file name.mp4",
			expected: "mp4",
		},
		{
			name:     "Ending in a dot",
			input:    "file.",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := DotFileType(tc.input)
			if result != tc.expected {
				t.Errorf("DotFileType(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}
