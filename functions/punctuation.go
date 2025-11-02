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
	
	// Handle multiple punctuation first (!!, ??, ...)
	multiplePunctPattern := regexp.MustCompile(`([!]{2,}|[?]{2,}|[.]{3,})([a-zA-Z'"\(])`)
	text = multiplePunctPattern.ReplaceAllString(text, "$1 $2")
	
	// Then handle single punctuation
	singlePunctPattern := regexp.MustCompile(`([.,:;!?])([a-zA-Z'"\(])`)
	text = singlePunctPattern.ReplaceAllString(text, "$1 $2")
	
	// Smart quote spacing: opening quotes need space before, closing quotes need space after
	text = p.handleQuoteSpacing(text)
	
	// Clean up multiple spaces
	multiSpacePattern := regexp.MustCompile(`\s+`)
	text = multiSpacePattern.ReplaceAllString(text, " ")
	
	return strings.TrimSpace(text)
}

// handleQuoteSpacing manages spacing around quotes based on whether they're opening or closing
func (p *PunctuationStage) handleQuoteSpacing(text string) string {
	result := []rune(text)
	singleQuoteOpen := true
	doubleQuoteOpen := true
	
	for i := 0; i < len(result); i++ {
		if result[i] == '\'' {
			if singleQuoteOpen {
				// Opening quote - ensure space before (but not at start)
				if i > 0 && result[i-1] != ' ' && result[i-1] != '\n' {
					// Insert space before opening quote
					result = append(result[:i], append([]rune{' '}, result[i:]...)...)
					i++
				}
			} else {
				// Closing quote - remove space before
				if i > 0 && result[i-1] == ' ' {
					// Remove space before closing quote
					result = append(result[:i-1], result[i:]...)
					i--
				}
			}
			singleQuoteOpen = !singleQuoteOpen
		} else if result[i] == '"' {
			if doubleQuoteOpen {
				// Opening quote - ensure space before (but not at start)
				if i > 0 && result[i-1] != ' ' && result[i-1] != '\n' {
					// Insert space before opening quote
					result = append(result[:i], append([]rune{' '}, result[i:]...)...)
					i++
				}
			} else {
				// Closing quote - remove space before
				if i > 0 && result[i-1] == ' ' {
					// Remove space before closing quote
					result = append(result[:i-1], result[i:]...)
					i--
				}
			}
			doubleQuoteOpen = !doubleQuoteOpen
		}
	}
	
	return string(result)
}