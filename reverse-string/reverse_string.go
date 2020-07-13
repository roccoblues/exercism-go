package reverse

import "strings"

func Reverse(input string) string {
	var b strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteString(string(runes[i]))
	}
	return b.String()
}
