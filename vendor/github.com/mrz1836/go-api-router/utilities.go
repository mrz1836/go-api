package apirouter

import (
	"regexp"
	"strings"
)

//camelCaseRe camel case regex
var camelCaseRe = regexp.MustCompile(`(?:^[\p{Ll}]|API|JSON|IP|_?\d+|_|[\p{Lu}]+)[\p{Ll}]*`)

//SnakeCase takes a camelCaseWord and breaks it into camel_case_word
func SnakeCase(str string) string {
	words := camelCaseRe.FindAllString(str, -1)

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(strings.Replace(words[i], "_", "", -1))
	}

	return strings.Join(words, "_")
}

//FindString returns the index of the first instance of needle in the array or -1 if it could not be found
func FindString(needle string, haystack []string) int {
	for i, straw := range haystack {
		if needle == straw {
			return i
		}
	}
	return -1
}
