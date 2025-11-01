package functions

import "testing"

func TestCaseTransform(t *testing.T) {
	stage := &CaseTransformStage{}

	tests := []struct {
		input    string
		expected string
	}{
		{"go (up)", "GO"},
		{"HELLO (low)", "hello"},
		{"bridge (cap)", "Bridge"},
		{"this is fun (up, 2)", "this IS FUN"},
		{"hello world (low, 2)", "hello world"},
		{"make this NICE (cap, 2)", "make This Nice"},
		{"test (cap, 1)", "Test"},
		{"a b c d (up, 3)", "a B C D"},
	}

	for _, test := range tests {
		result := stage.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}