package collatzconjecture

import "errors"

// CollatzConjecture calculates the number of steps needed for the given input
// to reach 1 in the Collatz Conjecture.
func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("not a positiv integer")
	}

	i := 0

	for input != 1 {
		i++
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
	}

	return i, nil
}
