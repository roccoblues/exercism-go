// Package isogram implements a methods to check wether a word is an isogram.
//
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.
package isogram

import "unicode"

// IsIsogram returns true if the given word is an isogram.
func IsIsogram(word string) bool {
	seen := map[rune]bool{}

	for _, l := range word {
		l = unicode.ToUpper(l)

		if l == ' ' || l == '-' {
			continue
		}

		if seen[l] {
			return false
		}

		seen[l] = true
	}

	return true
}
