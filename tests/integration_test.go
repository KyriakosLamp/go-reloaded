package main

import (
	"go-reloaded/functions"
	"testing"
)

func TestPhaseOneIntegration(t *testing.T) {
	// Create pipeline with first three stages: Numeric → Article → Case
	pipeline := NewPipeline()
	pipeline.AddStage(&functions.NumericConversionStage{})
	pipeline.AddStage(&functions.ArticleStage{})
	pipeline.AddStage(&functions.CaseTransformStage{})

	tests := []struct {
		input    string
		expected string
	}{
		{"1A (hex) files and 10 (bin) more (up, 2)", "26 files and 2 MORE"},
		{"FF (hex) is a amazing number (cap, 2)", "255 is an Amazing Number"},
		{"101 (bin) cats and a elephant (up)", "5 cats and an ELEPHANT"},
	}

	for _, test := range tests {
		result := pipeline.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}