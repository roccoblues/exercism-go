// Package rotationalcipher implements a rotational cipher
package rotationalcipher

// RotationalCipher shifts the input string by the given amount and returns it.
func RotationalCipher(input string, shift int) string {
	if shift == 0 {
		return input
	}

	output := make([]rune, len(input))

	for i, c := range input {
		switch {
		case c <= 'z' && c >= 'a':
			output[i] = ((c-'a')+rune(shift))%26 + 'a'
		case c <= 'Z' && c >= 'A':
			output[i] = ((c-'A')+rune(shift))%26 + 'A'
		default:
			output[i] = c
		}
	}

	return string(output)
}
