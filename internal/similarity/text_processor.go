package similarity

import (
	"regexp"
	"strings"
)

func PreProcess(text string) []string {
	// Convert to lowercase
	lowercaseText := strings.ToLower(text)

	// Remove punctuation using regular expression
	// This regex removes all non-word characters
	reg := regexp.MustCompile(`[^\w\s]`)
	clearedText := reg.ReplaceAllString(lowercaseText, "")

	// Split into words and remove any empty strings
	words := strings.Fields(clearedText)

	return words
}
