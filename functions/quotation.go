package functions

import (
	"regexp"
)

// QuotationStage handles quotation mark spacing and formatting
type QuotationStage struct{}

// Process removes inner spaces from quoted text, making quotes hug the content
func (q *QuotationStage) Process(text string) string {
	// Compile trim pattern once
	trimPattern := regexp.MustCompile(`^\s+|\s+$`)
	
	// Handle single quotes with content that may include punctuation
	singleQuotePattern := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	text = singleQuotePattern.ReplaceAllStringFunc(text, func(match string) string {
		content := match[1 : len(match)-1] // Remove outer quotes
		content = trimPattern.ReplaceAllString(content, "") // Trim spaces
		return "'" + content + "'"
	})
	
	// Handle double quotes with content that may include punctuation
	doubleQuotePattern := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
	text = doubleQuotePattern.ReplaceAllStringFunc(text, func(match string) string {
		content := match[1 : len(match)-1] // Remove outer quotes
		content = trimPattern.ReplaceAllString(content, "") // Trim spaces
		return "\"" + content + "\""
	})
	
	return text
}