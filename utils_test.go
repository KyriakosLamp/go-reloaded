package main

import (
	"os"
	"testing"
	"go-reloaded/functions"
)

func TestFileIO(t *testing.T) {
	// Test data
	testContent := "Hello, World!"
	testFile := "test_file.txt"

	// Clean up after test
	defer os.Remove(testFile)

	// Test WriteFile
	err := functions.WriteFile(testFile, testContent)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	// Test ReadFile
	content, err := functions.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	// Verify content matches
	if content != testContent {
		t.Errorf("Expected %q, got %q", testContent, content)
	}
}