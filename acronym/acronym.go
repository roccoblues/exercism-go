package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate convert a phrase to its acronym.
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, split)

	acronym := make([]byte, len(words))

	for i, w := range words {
		acronym[i] = w[0]
	}

	return strings.ToUpper(string(acronym))
}

func split(c rune) bool {
	return unicode.IsSpace(c) || c == '-'
}
