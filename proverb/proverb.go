package proverb

import "fmt"

// Proverb generates a rhyme from the given input words.
func Proverb(words []string) []string {
	rhyme := make([]string, len(words))

	if len(words) == 0 {
		return rhyme
	}

	for i, word := range words[:len(words)-1] {
		rhyme[i] = fmt.Sprintf("For want of a %s the %s was lost.", word, words[i+1])
	}

	rhyme[len(words)-1] = fmt.Sprintf("And all for the want of a %s.", words[0])

	return rhyme
}
