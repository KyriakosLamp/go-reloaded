package functions

import (
	"regexp"
)

// QuotationStage handles quotation mark spacing and formatting
type QuotationStage struct{}

// Process removes inner spaces from quoted text, making quotes hug the content
func (q *QuotationStage) Process(text string) string {
	// Handle single quotes
	singleQuotePattern := regexp.MustCompile(`'\s+([^']*?)\s+'`)
	text = singleQuotePattern.ReplaceAllString(text, "'$1'")
	
	// Handle double quotes
	doubleQuotePattern := regexp.MustCompile(`"\s+([^"]*?)\s+"`)
	text = doubleQuotePattern.ReplaceAllString(text, "\"$1\"")
	
	return text
}