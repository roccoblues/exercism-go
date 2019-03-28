// Package grains implements methods related to the grains of wheat on a
// chessboard problem. (https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem)
package grains

import "errors"

// Square calculates how many grains are on the given square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be between 1 and 64")
	}

	return 1 << uint(n-1), nil
}

// Total calculates the total amount of grains on the chessboard.
// Can be solved easily with: (2^64)-1
func Total() uint64 {
	return 1<<64 - 1
}
