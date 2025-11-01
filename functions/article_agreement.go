package functions

import (
	"regexp"
	"strings"
)

// ArticleStage handles article agreement transformation (a -> an)
type ArticleStage struct{}

// Process replaces "a" with "an" before words starting with vowels (a,e,i,o,u) or 'h'
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