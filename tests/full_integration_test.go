package main

import (
	"go-reloaded/functions"
	"testing"
)

func TestFullPipelineIntegration(t *testing.T) {
	// Create full pipeline: NumericConversion → Article → Case → Quotation → Punctuation
	pipeline := NewPipeline()
	pipeline.AddStage(&functions.NumericConversionStage{})
	pipeline.AddStage(&functions.ArticleStage{})
	pipeline.AddStage(&functions.CaseTransformStage{})
	pipeline.AddStage(&functions.QuotationStage{})
	pipeline.AddStage(&functions.PunctuationStage{})

	// Formula One test case
	input := `it (cap) was a historic race ,lewis hamilton (up,2) took pole position  , scoring " 1E (hex) points  " . it (cap) was ' amazing ' to watch a incredible battle on track ! The (low) crowd shouted (up) : go lewis go (cap, 3) !! the fans waved flags (cap, 4) and a HERO (low) returned to the podium . WHAT A INCREDIBLE MOMENT (low, 3) for the sport !`
	
	expected := `It was an historic race, LEWIS HAMILTON took pole position, scoring "30 points". It was 'amazing' to watch an incredible battle on track! the crowd SHOUTED: Go Lewis Go!! The Fans Waved Flags and an hero returned to the podium. WHAT a incredible moment for the sport!`

	result := pipeline.Process(input)
	
	if result != expected {
		t.Errorf("Formula One test failed.\nInput: %q\nExpected: %q\nGot: %q", input, expected, result)
	}
}

func TestComplexCombinations(t *testing.T) {
	// Create full pipeline
	pipeline := NewPipeline()
	pipeline.AddStage(&functions.NumericConversionStage{})
	pipeline.AddStage(&functions.ArticleStage{})
	pipeline.AddStage(&functions.CaseTransformStage{})
	pipeline.AddStage(&functions.QuotationStage{})
	pipeline.AddStage(&functions.PunctuationStage{})

	tests := []struct {
		input    string
		expected string
		name     string
	}{
		{"1A (hex) files and 10 (bin) more (up)", "26 files and 2 MORE", "Numeric + Case"},
		{"He shouted (up, 2): ' run now ' !!", "HE SHOUTED: 'run now'!!", "Case + Quotes + Punctuation"},
		{"I met a (cap) honest man.", "I met A honest man.", "Article + Case"},
		{"a old hero (cap, 3)", "An Old Hero", "Article + Case Multi-word"},
		{"wow (up) ... great ,job !", "WOW... great, job!", "Case + Punctuation"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pipeline.Process(test.input)
			if result != test.expected {
				t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
			}
		})
	}
}