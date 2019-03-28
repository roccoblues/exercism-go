// Package twofer contains the solution to the "Two Fer" exercise.
package twofer

import "fmt"

// ShareWith returns the "two for one string".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
