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
	// Pattern to match transformation marker and capture only nearby preceding words
	pattern := regexp.MustCompile(`((?:\S+[ \t]+){1,5}?)\((up|low|cap)(?:,\s*(\d+))?\)`)
	
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
		
		// Find actual word positions (not punctuation-only)
		wordPositions := []int{}
		for i, word := range words {
			// Check if word contains at least one letter or digit
			hasLetterOrDigit := false
			for _, r := range word {
				if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
					hasLetterOrDigit = true
					break
				}
			}
			if hasLetterOrDigit {
				wordPositions = append(wordPositions, i)
			}
		}
		
		if count > len(wordPositions) {
			count = len(wordPositions)
		}
		
		// Apply transformation to the last 'count' actual words
		for i := len(wordPositions) - count; i < len(wordPositions); i++ {
			wordIndex := wordPositions[i]
			switch transform {
			case "up":
				words[wordIndex] = strings.ToUpper(words[wordIndex])
			case "low":
				words[wordIndex] = strings.ToLower(words[wordIndex])
			case "cap":
				if len(words[wordIndex]) > 0 {
					words[wordIndex] = strings.ToUpper(string(words[wordIndex][0])) + strings.ToLower(words[wordIndex][1:])
				}
			}
		}
		
		// Replace the matched part with transformed words
		transformed := strings.Join(words, " ")
		text = strings.Replace(text, match[0], transformed, 1)
	}
	
	return text
}