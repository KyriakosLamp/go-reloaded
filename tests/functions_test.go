package main

import (
	"go-reloaded/functions"
	"testing"
)

func TestNumericConversion(t *testing.T) {
	stage := &functions.NumericConversionStage{}

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

func TestArticleAgreement(t *testing.T) {
	stage := &functions.ArticleStage{}

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

func TestCaseTransform(t *testing.T) {
	stage := &functions.CaseTransformStage{}

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

func TestQuotation(t *testing.T) {
	stage := &functions.QuotationStage{}

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

func TestPunctuation(t *testing.T) {
	stage := &functions.PunctuationStage{}

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