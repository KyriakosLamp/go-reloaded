package functions

import "testing"

func TestNumericConversion(t *testing.T) {
	stage := &NumericConversionStage{}

	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex)", "30"},
		{"10 (bin)", "2"},
		{"1A (hex) and 101 (bin)", "26 and 5"},
		{"FF (hex) files", "255 files"},
		{"1010 (bin) days", "10 days"},
		{"XYZ (hex)", "XYZ (hex)"}, // Invalid hex should remain unchanged
		{"102 (bin)", "102 (bin)"},  // Invalid binary should remain unchanged
	}

	for _, test := range tests {
		result := stage.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}