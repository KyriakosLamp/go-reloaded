package functions

import (
	"regexp"
)

// QuotationStage handles quotation mark formatting
type QuotationStage struct{}

// Process removes inner spacing from quoted text
func (q *QuotationStage) Process(text string) string {
	// Handle single quotes
	singleQuotePattern := regexp.MustCompile(`'\s+([^']*?)\s+'`)
	text = singleQuotePattern.ReplaceAllString(text, "'$1'")
	
	// Handle double quotes
	doubleQuotePattern := regexp.MustCompile(`"\s+([^"]*?)\s+"`)
	text = doubleQuotePattern.ReplaceAllString(text, "\"$1\"")
	
	return text
}