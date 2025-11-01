package functions

import "testing"

func TestArticleAgreement(t *testing.T) {
	stage := &ArticleStage{}

	tests := []struct {
		input    string
		expected string
	}{
		{"a orange car", "an orange car"},
		{"a honest man", "an honest man"},
		{"a user", "an user"},
		{"a amazing car", "an amazing car"},
		{"a elephant", "an elephant"},
		{"a incredible battle", "an incredible battle"},
		{"a hero", "an hero"},
		{"a book", "a book"},     // Should not change
		{"a cat", "a cat"},       // Should not change
		{"I met a honest person", "I met an honest person"},
	}

	for _, test := range tests {
		result := stage.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}