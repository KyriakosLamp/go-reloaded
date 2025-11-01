package functions

import (
	"regexp"
	"strings"
)

// PunctuationStage handles punctuation spacing and grouping
type PunctuationStage struct{}

// Process fixes punctuation spacing - attach to previous word, space after
func (p *PunctuationStage) Process(text string) string {
	// Pattern to match spaces before punctuation marks
	pattern := regexp.MustCompile(`\s+([.,:;!?]+)`)
	
	// Remove spaces before punctuation
	text = pattern.ReplaceAllString(text, "$1")
	
	// Pattern to ensure space after punctuation (but not at end of text)
	spaceAfterPattern := regexp.MustCompile(`([.,:;!?]+)([^\s.,:;!?])`)
	text = spaceAfterPattern.ReplaceAllString(text, "$1 $2")
	
	// Clean up multiple spaces
	multiSpacePattern := regexp.MustCompile(`\s+`)
	text = multiSpacePattern.ReplaceAllString(text, " ")
	
	return strings.TrimSpace(text)
}