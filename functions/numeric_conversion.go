package functions

import (
	"regexp"
	"strconv"
	"strings"
)

// NumericConversionStage handles hex and binary conversions
type NumericConversionStage struct{}

// Process converts hex and binary numbers to decimal
func (n *NumericConversionStage) Process(text string) string {
	// Pattern to match (hex) or (bin) markers
	hexPattern := regexp.MustCompile(`(\w+)\s+\(hex\)`)
	binPattern := regexp.MustCompile(`(\w+)\s+\(bin\)`)

	// Process hex conversions
	text = hexPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) >= 2 {
			hexStr := parts[0]
			if decimal, err := strconv.ParseInt(hexStr, 16, 64); err == nil {
				return strconv.FormatInt(decimal, 10)
			}
		}
		return match // Return original if conversion fails
	})

	// Process binary conversions
	text = binPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) >= 2 {
			binStr := parts[0]
			if decimal, err := strconv.ParseInt(binStr, 2, 64); err == nil {
				return strconv.FormatInt(decimal, 10)
			}
		}
		return match // Return original if conversion fails
	})

	return text
}