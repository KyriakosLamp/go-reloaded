package functions

import "testing"

func TestPunctuation(t *testing.T) {
	stage := &PunctuationStage{}

	tests := []struct {
		input    string
		expected string
	}{
		{"Hello ,world !!", "Hello, world!!"},
		{"Wait ... what ?", "Wait... what?"},
		{"Hello , world !", "Hello, world!"},
		{"Great job .", "Great job."},
		{"What ?! Really ...", "What?! Really..."},
		{"No punctuation here", "No punctuation here"}, // Should not change
		{"End of sentence .", "End of sentence."},
		{"Multiple   spaces", "Multiple spaces"}, // Clean up spaces
	}

	for _, test := range tests {
		result := stage.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}