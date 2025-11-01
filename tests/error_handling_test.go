package main

import (
	"go-reloaded/functions"
	"testing"
)

func TestErrorHandling(t *testing.T) {
	// Enable logging for testing
	functions.GlobalLogger = functions.NewLogger(true)
	defer func() {
		// Disable logging after test
		functions.GlobalLogger = functions.NewLogger(false)
	}()

	stage := &functions.NumericConversionStage{}

	tests := []struct {
		input    string
		expected string
		name     string
	}{
		{"XYZ (hex)", "XYZ (hex)", "Invalid hex should remain unchanged"},
		{"102 (bin)", "102 (bin)", "Invalid binary should remain unchanged"},
		{"GHI (hex) and 123 (bin)", "GHI (hex) and 123 (bin)", "Multiple invalid conversions"},
		{"1A (hex) and XYZ (hex)", "26 and XYZ (hex)", "Mixed valid and invalid hex"},
		{"10 (bin) and 102 (bin)", "2 and 102 (bin)", "Mixed valid and invalid binary"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stage.Process(test.input)
			if result != test.expected {
				t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
			}
		})
	}
}

func TestLoggerFunctionality(t *testing.T) {
	// Test logger creation and basic functionality
	logger := functions.NewLogger(true)
	
	// These should not crash the program
	logger.LogWarning("Test warning message")
	logger.LogError("Test error message", false) // Non-critical error
	
	// Test disabled logger
	disabledLogger := functions.NewLogger(false)
	disabledLogger.LogWarning("This should not appear")
	disabledLogger.LogError("This should not appear", false)
}