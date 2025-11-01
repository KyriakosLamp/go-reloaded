package functions

import (
	"regexp"
	"strconv"
	"strings"
)

// CaseTransformStage handles case transformations with optional word count
type CaseTransformStage struct{}

// Process applies (up), (low), (cap) transformations with optional count parameter
func (c *CaseTransformStage) Process(text string) string {
	// Pattern to match word(s) followed by transformation marker
	pattern := regexp.MustCompile(`(\S+(?:\s+\S+)*?)\s+\((up|low|cap)(?:,\s*(\d+))?\)`)
	
	for {
		match := pattern.FindStringSubmatch(text)
		if match == nil {
			break
		}
		
		wordsPart := match[1]
		transform := match[2]
		count := 1
		if len(match) > 3 && match[3] != "" {
			if n, err := strconv.Atoi(match[3]); err == nil {
				count = n
			}
		}
		
		words := strings.Fields(wordsPart)
		if count > len(words) {
			count = len(words)
		}
		
		// Apply transformation to the last 'count' words
		for i := len(words) - count; i < len(words); i++ {
			switch transform {
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "low":
				words[i] = strings.ToLower(words[i])
			case "cap":
				if len(words[i]) > 0 {
					words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
				}
			}
		}
		
		// Replace the matched part with transformed words
		transformed := strings.Join(words, " ")
		text = strings.Replace(text, match[0], transformed, 1)
	}
	
	return text
}