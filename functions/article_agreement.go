package functions

import (
	"regexp"
	"strings"
)

// ArticleStage handles article agreement (a -> an)
type ArticleStage struct{}

// Process replaces "a" with "an" before vowels and 'h'
func (a *ArticleStage) Process(text string) string {
	// Pattern to match "a" followed by a word starting with vowel or h
	pattern := regexp.MustCompile(`\ba\s+([aeiouAEIOUhH]\w*)`)
	
	return pattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			return "an " + parts[1]
		}
		return match
	})
}