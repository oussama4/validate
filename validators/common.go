package validators

import (
	"strings"
	"unicode"
)

var CustomKeys = map[string]string{}

func ToUnderscore(s string) string {
	// Convert to lowercase first
	s = strings.ToLower(s)

	var result strings.Builder
	var prevChar rune

	for i, char := range s {
		// Handle special characters and spaces
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			if prevChar != '_' {
				result.WriteRune('_')
			}
			prevChar = '_'
			continue
		}

		// Handle camelCase: insert underscore before uppercase letters
		if i > 0 && unicode.IsUpper(char) && prevChar != '_' {
			result.WriteRune('_')
		}

		// Write the current character
		result.WriteRune(char)
		prevChar = char
	}

	// Trim leading and trailing underscores
	return strings.Trim(result.String(), "_")
}
func GenerateKey(s string) string {
	key := CustomKeys[s]
	if key != "" {
		return key
	}
	return ToUnderscore(s)
}
