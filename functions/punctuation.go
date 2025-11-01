package functions

import (
	"regexp"
	"strings"
)

// PunctuationStage handles punctuation spacing and grouping normalization
type PunctuationStage struct{}

// Process attaches punctuation to preceding words and ensures proper spacing
func (p *PunctuationStage) Process(text string) string {
	// Pattern to match spaces before punctuation marks (but not inside quotes)
	pattern := regexp.MustCompile(`\s+([.,:;!?]+)`)
	
	// Remove spaces before punctuation
	text = pattern.ReplaceAllString(text, "$1")
	
	// Special case: ensure space after colon when followed by quote
	colonQuotePattern := regexp.MustCompile(`(:)(['"]+)`)
	text = colonQuotePattern.ReplaceAllString(text, "$1 $2")
	
	// Pattern to ensure space after punctuation (but not at end of text or before quotes)
	spaceAfterPattern := regexp.MustCompile(`([.,:;!?]+)([^\s.,:;!?'"\n])`)
	text = spaceAfterPattern.ReplaceAllString(text, "$1 $2")
	
	// Remove spaces between punctuation and closing quotes (except colon and comma)
	punctQuotePattern := regexp.MustCompile(`([.;!?])\s+(['"]+)`)
	text = punctQuotePattern.ReplaceAllString(text, "$1$2")
	
	// Clean up multiple spaces
	multiSpacePattern := regexp.MustCompile(`\s+`)
	text = multiSpacePattern.ReplaceAllString(text, " ")
	
	return strings.TrimSpace(text)
}