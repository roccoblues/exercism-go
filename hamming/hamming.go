// Package hamming implements functions to calculate the hamming distance.
package hamming

import "errors"

// Distance calculates the hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("different input length")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
