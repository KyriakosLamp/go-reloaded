package functions

import (
	"regexp"
	"strings"
)

// ArticleStage handles article agreement transformation (a -> an)
type ArticleStage struct{}

// Process replaces "a" with "an" and handles case transformation patterns
func (a *ArticleStage) Process(text string) string {
	// Pattern to match "a" followed by case transformation and vowel-starting word
	caseTransformPattern := regexp.MustCompile(`\ba\s+\((cap|up|low)\)\s+([aeiouAEIOUhH]\w*)`)
	text = caseTransformPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := caseTransformPattern.FindStringSubmatch(match)
		if len(parts) == 3 {
			transformType := parts[1]
			word := parts[2]
			switch transformType {
			case "cap":
				return "An (cap) " + word
			case "up":
				return "AN (up) " + word
			case "low":
				return "an (low) " + word
			}
		}
		return match
	})
	
	// Pattern to match regular "a" followed by a word starting with vowel or h
	lowerPattern := regexp.MustCompile(`\ba\s+([aeiouAEIOUhH]\w*)`)
	text = lowerPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			return "an " + parts[1]
		}
		return match
	})
	
	// Pattern to match uppercase "A" followed by a word starting with vowel or h
	upperPattern := regexp.MustCompile(`\bA\s+([aeiouAEIOUhH]\w*)`)
	text = upperPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		if len(parts) == 2 {
			return "AN " + parts[1]
		}
		return match
	})
	
	return text
}