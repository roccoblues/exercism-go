package bob

import (
	"strings"
	"unicode"
)

// Hey answers the given remark.
func Hey(remark string) string {
	s := strings.TrimSpace(remark)

	if len(s) == 0 {
		return "Fine. Be that way!"
	}

	if hasLetters(s) && isCaps(s) && isQuestion(s) {
		return "Calm down, I know what I'm doing!"
	}

	if hasLetters(s) && isCaps(s) {
		return "Whoa, chill out!"
	}

	if isQuestion(s) {
		return "Sure."
	}

	return "Whatever."
}

func hasLetters(s string) bool {
	return strings.IndexFunc(s, unicode.IsLetter) >= 0
}

func isCaps(s string) bool {
	return strings.ToUpper(s) == s
}

func isQuestion(s string) bool {
	return s[len(s)-1] == '?'
}
