// Package scrabble implements a method to calculate the scrabble score for a word.
package scrabble

// Score calculates the scrabble score for the given word.
func Score(word string) int {
	score := 0

	for _, l := range word {
		switch l {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T', 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'D', 'G', 'd', 'g':
			score += 2
		case 'B', 'C', 'M', 'P', 'b', 'c', 'm', 'p':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y', 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'K', 'k':
			score += 5
		case 'J', 'X', 'j', 'x':
			score += 8
		case 'Q', 'Z', 'q', 'z':
			score += 10
		}
	}

	return score
}
