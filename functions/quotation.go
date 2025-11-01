package functions

import (
	"regexp"
)

// QuotationStage handles quotation mark spacing and formatting
type QuotationStage struct{}

// Process removes inner spaces from quoted text, making quotes hug the content
func (q *QuotationStage) Process(text string) string {
	// Handle single quotes with content that may include punctuation
	// This pattern captures everything between quotes and trims internal spaces
	singleQuotePattern := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	text = singleQuotePattern.ReplaceAllStringFunc(text, func(match string) string {
		// Extract content between quotes
		content := match[1 : len(match)-1] // Remove outer quotes
		content = regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(content, "") // Trim spaces
		return "'" + content + "'"
	})
	
	// Handle double quotes with content that may include punctuation
	doubleQuotePattern := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	text = doubleQuotePattern.ReplaceAllStringFunc(text, func(match string) string {
		// Extract content between quotes
		content := match[1 : len(match)-1] // Remove outer quotes
		content = regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(content, "") // Trim spaces
		return "\"" + content + "\""
	})
	
	return text
}