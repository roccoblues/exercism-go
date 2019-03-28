// Package luhn implement a method to if a number is valid per the Luhn formula.
// https://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

import (
	"unicode"
)

// Valid checks if the input is valid per the Luhn formula.
func Valid(input string) bool {
	sum := 0
	p := 0

	for i := len(input) - 1; i >= 0; i-- {
		r := rune(input[i])

		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsNumber(r) {
			return false
		}

		n := int(r) - '0'
		p++

		if p%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
	}

	if p <= 1 {
		return false
	}

	return sum%10 == 0
}
