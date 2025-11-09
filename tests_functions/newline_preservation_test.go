package main

import (
	"go-reloaded/functions"
	"testing"
)

func TestNewlinePreservation(t *testing.T) {
	// Test that newlines are preserved through the pipeline
	pipeline := NewPipeline()
	pipeline.AddStage(&functions.NumericConversionStage{})
	pipeline.AddStage(&functions.PunctuationStage{})
	pipeline.AddStage(&functions.ArticleStage{})
	pipeline.AddStage(&functions.CaseTransformStage{})
	pipeline.AddStage(&functions.QuotationStage{})

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Multi-line with transformations",
			input: `1E (hex)
SHOUTING (low)
welcome (cap)
A amazing day
Hello ,world !`,
			expected: `30
shouting
Welcome
An amazing day
Hello, world!`,
		},
		{
			name: "Single line should not add newlines",
			input: "hello (cap) world",
			expected: "Hello world",
		},
		{
			name: "Empty lines preserved",
			input: `first (cap)

second (up)`,
			expected: `First

SECOND`,
		},
		{
			name: "Complex multi-line with all transformations",
			input: `this (cap) is a awesome day
101 (bin) more days
' hello world '`,
			expected: `This is an awesome day
5 more days
'hello world'`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := pipeline.Process(test.input)
			if result != test.expected {
				t.Errorf("Expected:\n%q\nGot:\n%q", test.expected, result)
			}
		})
	}
}

func TestIndividualStageNewlinePreservation(t *testing.T) {
	tests := []struct {
		name     string
		stage    Stage
		input    string
		expected string
	}{
		{
			name:  "NumericConversion preserves newlines",
			stage: &functions.NumericConversionStage{},
			input: `1E (hex)
10 (bin)`,
			expected: `30
2`,
		},
		{
			name:  "Punctuation preserves newlines",
			stage: &functions.PunctuationStage{},
			input: `Hello ,world !
Great job .`,
			expected: `Hello, world!
Great job.`,
		},
		{
			name:  "Article preserves newlines",
			stage: &functions.ArticleStage{},
			input: `a apple
a orange`,
			expected: `an apple
an orange`,
		},
		{
			name:  "Case transform preserves newlines",
			stage: &functions.CaseTransformStage{},
			input: `hello (cap)
world (up)`,
			expected: `Hello
WORLD`,
		},
		{
			name:  "Quotation preserves newlines",
			stage: &functions.QuotationStage{},
			input: `' hello '
" world "`,
			expected: `'hello'
"world"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.stage.Process(test.input)
			if result != test.expected {
				t.Errorf("Expected:\n%q\nGot:\n%q", test.expected, result)
			}
		})
	}
}