package functions

import (
	"regexp"
	"strings"
)

// PunctuationStage handles punctuation spacing and grouping normalization
type PunctuationStage struct{}

// Process attaches punctuation to preceding words and ensures proper spacing
func (p *PunctuationStage) Process(text string) string {
	// Pattern to match spaces before punctuation marks (but not inside quotes, within same line)
	pattern := regexp.MustCompile(`[ \t]+([.,:;!?]+)`)
	
	// Remove spaces before punctuation
	text = pattern.ReplaceAllString(text, "$1")
	
	// Special case: ensure space after colon when followed by quote
	colonQuotePattern := regexp.MustCompile(`(:)(['"]+)`)
	text = colonQuotePattern.ReplaceAllString(text, "$1 $2")
	
	// Handle multiple punctuation first (!!, ??, ...)
	multiplePunctPattern := regexp.MustCompile(`([!?]{2,}|[.]{3,})([a-zA-Z'"(])`)
	text = multiplePunctPattern.ReplaceAllString(text, "$1 $2")
	
	// Then handle single punctuation
	singlePunctPattern := regexp.MustCompile(`([.,:;!?])([a-zA-Z'"(])`)
	text = singlePunctPattern.ReplaceAllString(text, "$1 $2")
	
	// Smart quote spacing: opening quotes need space before, closing quotes need space after
	text = p.handleQuoteSpacing(text)
	
	// Clean up multiple spaces but preserve newlines
	multiSpacePattern := regexp.MustCompile(`[ \t]+`)
	text = multiSpacePattern.ReplaceAllString(text, " ")
	
	// Trim spaces from each line but preserve newlines
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}

// handleQuoteSpacing manages spacing around quotes based on whether they're opening or closing
func (p *PunctuationStage) handleQuoteSpacing(text string) string {
	result := []rune(text)
	singleQuoteOpen := true
	doubleQuoteOpen := true
	
	for i := 0; i < len(result); i++ {
		if result[i] == '\'' {
			// Check if this is a consecutive quote (don't add space between consecutive quotes)
			isConsecutive := (i > 0 && result[i-1] == '\'') || (i < len(result)-1 && result[i+1] == '\'')
			
			if singleQuoteOpen {
				// Opening quote - ensure space before (but not at start or if consecutive)
				if i > 0 && result[i-1] != ' ' && result[i-1] != '\n' && result[i-1] != '\'' {
					// Insert space before opening quote
					result = append(result[:i], append([]rune{' '}, result[i:]...)...)
					i++
				}
			} else {
				// Closing quote - remove space before (but not if consecutive)
				if i > 0 && result[i-1] == ' ' && !isConsecutive {
					// Remove space before closing quote
					result = append(result[:i-1], result[i:]...)
					i--
				}
			}
			// Only flip state if not consecutive (consecutive quotes don't change open/close state)
			if !isConsecutive {
				singleQuoteOpen = !singleQuoteOpen
			}
		} else if result[i] == '"' {
			// Check if this is a consecutive quote
			isConsecutive := (i > 0 && result[i-1] == '"') || (i < len(result)-1 && result[i+1] == '"')
			
			if doubleQuoteOpen {
				// Opening quote - ensure space before (but not at start or if consecutive)
				if i > 0 && result[i-1] != ' ' && result[i-1] != '\n' && result[i-1] != '"' {
					// Insert space before opening quote
					result = append(result[:i], append([]rune{' '}, result[i:]...)...)
					i++
				}
			} else {
				// Closing quote - remove space before (but not if consecutive)
				if i > 0 && result[i-1] == ' ' && !isConsecutive {
					// Remove space before closing quote
					result = append(result[:i-1], result[i:]...)
					i--
				}
			}
			// Only flip state if not consecutive
			if !isConsecutive {
				doubleQuoteOpen = !doubleQuoteOpen
			}
		}
	}
	
	return string(result)
}