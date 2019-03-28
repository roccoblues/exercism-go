// Package allergies implements methods to test for allergies.
package allergies

var substances = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies returns the list of allergies for the given score.
func Allergies(score uint) []string {
	allergies := []string{}

	for i, v := range substances {
		if score&(1<<uint(i)) != 0 {
			allergies = append(allergies, v)
		}
	}

	return allergies
}

// AllergicTo returns true if someone is allergic to the given substance.
func AllergicTo(score uint, substance string) bool {
	for i, v := range substances {
		if v == substance {
			return score&(1<<uint(i)) != 0
		}
	}

	return false
}
