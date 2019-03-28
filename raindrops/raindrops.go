// Package raindrops implements the solution to the Raindrops exercise.
package raindrops

import "strconv"

// Convert converts a number to a string, based on the number's factors.
func Convert(input int) string {
	var output string

	if input%3 == 0 {
		output = output + "Pling"
	}
	if input%5 == 0 {
		output = output + "Plang"
	}
	if input%7 == 0 {
		output = output + "Plong"
	}

	if output == "" {
		output = strconv.Itoa(input)
	}

	return output
}
