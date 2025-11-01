package main

import (
	"go-reloaded/functions"
	"testing"
)

func TestPhaseTwoIntegration(t *testing.T) {
	// Create pipeline with Quotation → Punctuation stages
	pipeline := NewPipeline()
	pipeline.AddStage(&functions.QuotationStage{})
	pipeline.AddStage(&functions.PunctuationStage{})

	tests := []struct {
		input    string
		expected string
	}{
		{"He said ' hello world ' ,and smiled .", "He said 'hello world', and smiled."},
		{"She replied \" amazing \" !!", "She replied \"amazing\"!!"},
		{"Wait ... ' what ' ?", "Wait... 'what'?"},
		{"The \" quick brown \" fox ,jumped .", "The \"quick brown\" fox, jumped."},
	}

	for _, test := range tests {
		result := pipeline.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}