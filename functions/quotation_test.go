package functions

import "testing"

func TestQuotation(t *testing.T) {
	stage := &QuotationStage{}

	tests := []struct {
		input    string
		expected string
	}{
		{"He said: ' wow '.", "He said: 'wow'."},
		{"\" hello \"", "\"hello\""},
		{"' awesome '", "'awesome'"},
		{"\" amazing \"", "\"amazing\""},
		{"He said ' hello world ' to me.", "He said 'hello world' to me."},
		{"The \" quick brown \" fox.", "The \"quick brown\" fox."},
		{"No quotes here", "No quotes here"}, // Should not change
		{"' single ' and \" double \"", "'single' and \"double\""},
	}

	for _, test := range tests {
		result := stage.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}