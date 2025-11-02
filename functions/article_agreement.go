package functions

import (
	"regexp"
	"strings"
)

// ArticleStage handles article agreement transformation (a -> an)
type ArticleStage struct{}

// Process replaces "a"/"A" with "an"/"An" before vowel-starting words, handling keywords properly
func (a *ArticleStage) Process(text string) string {
	// Pattern to match article + space + next word (including transformation keywords)
	pattern := regexp.MustCompile(`\b([aA])\s+(?:\([a-z]+(?:,\s*\d+)?\)\s+)*([aeiouAEIOUhH]\w*)`)
	
	// Replace all matches
	text = pattern.ReplaceAllStringFunc(text, func(match string) string {
		// Find the article at the beginning
		if match[0] == 'a' {
			return strings.Replace(match, "a ", "an ", 1)
		} else if match[0] == 'A' {
			return strings.Replace(match, "A ", "An ", 1)
		}
		return match
	})
	
	return text
}