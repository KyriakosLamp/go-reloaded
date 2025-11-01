package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// NumericConversionStage handles conversion of hexadecimal and binary numbers to decimal
type NumericConversionStage struct{}

// Process converts (hex) and (bin) markers with preceding numbers to decimal format
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
			} else {
				GlobalLogger.LogWarning(fmt.Sprintf("Invalid hex number: %s", hexStr))
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
			} else {
				GlobalLogger.LogWarning(fmt.Sprintf("Invalid binary number: %s", binStr))
			}
		}
		return match // Return original if conversion fails
	})

	return text
}