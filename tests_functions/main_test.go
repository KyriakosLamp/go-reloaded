package main

import (
	"os"
	"testing"
	"go-reloaded/functions"
)

func TestCLIIntegration(t *testing.T) {
	// Create test input file
	inputFile := "test_input.txt"
	outputFile := "test_output.txt"
	testContent := "Hello World"

	// Clean up after test
	defer os.Remove(inputFile)
	defer os.Remove(outputFile)

	// Create input file
	err := functions.WriteFile(inputFile, testContent)
	if err != nil {
		t.Fatalf("Failed to create test input file: %v", err)
	}

	// Create pipeline and process
	pipeline := NewPipeline()
	result := pipeline.Process(testContent)

	// Write output
	err = functions.WriteFile(outputFile, result)
	if err != nil {
		t.Fatalf("Failed to write output file: %v", err)
	}

	// Verify output file exists and has correct content
	content, err := functions.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if content != testContent {
		t.Errorf("Expected %q, got %q", testContent, content)
	}
}

func TestPipelineStructure(t *testing.T) {
	pipeline := NewPipeline()
	
	// Test empty pipeline
	input := "test text"
	result := pipeline.Process(input)
	
	if result != input {
		t.Errorf("Empty pipeline should return unchanged text, got %q", result)
	}
}