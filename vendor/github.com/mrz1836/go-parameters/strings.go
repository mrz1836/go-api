package parameters

import (
	"regexp"
	"strings"
	"unicode"
)

// KnownAbbreviations contains lower case versions of abbreviations to match.
// Any entry in this list will become full upper case when converting from
// snake_case to camelCase
//
// 		user_id -> UserID
var KnownAbbreviations = []string{"id", "json", "html", "xml"}

var camelCaseRe = regexp.MustCompile(`(?:^[\p{Ll}]|\d+|[\p{Lu}]+)[\p{Ll}]*`)

// CamelToSnakeCase converts CamelCase to snake_case
// Consecutive capital letters will be treated as one word:
//  HTML -> html
func CamelToSnakeCase(str string) string {
	words := camelCaseRe.FindAllString(str, -1)

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	return strings.Join(words, "_")
}

// SnakeToCamelCase converts snake_case to CamelCase.
// When:
//  ucFirst = false - snake_case -> snakeCase
//  ucFirst = true  - snake_case -> SnakeCase
func SnakeToCamelCase(str string, ucFirst bool) string {
	words := strings.Split(str, "_")
	var i int
	if ucFirst {
		i = 0
	} else {
		i = 1
	}

	for ; i < len(words); i++ {
		if isKnownAbbreviation(words[i]) {
			words[i] = strings.ToUpper(words[i])
		} else {
			words[i] = MakeFirstUpperCase(words[i])
		}
	}

	return strings.Join(words, "")
}

// MakeFirstUpperCase upper cases the first letter of the string
func MakeFirstUpperCase(s string) string {

	// Handle empty and 1 character strings
	if len(s) < 2 {
		return strings.ToUpper(s)
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// isKnownAbbreviation is a known abbreviation
func isKnownAbbreviation(word string) bool {
	word = strings.ToLower(word)

	for _, value := range KnownAbbreviations {
		if value == word {
			return true
		}
	}

	return false
}
