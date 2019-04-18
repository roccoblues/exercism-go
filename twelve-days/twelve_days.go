// Package twelve contains methods to ouput the lyrics to 'The Twelve Days of Christmas'.
package twelve

import "fmt"

var gifts = [...]string{
	"a Partridge",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var days = [...]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

// Song returns the lyrics to 'The Twelve Days of Christmas'.
func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}

	return song
}

// Verse returns specific verse of 'The Twelve Days of Christmas'.
func Verse(day int) string {
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[day-1])

	if day > 1 {
		for i := day; i > 1; i-- {
			verse += gifts[i-1] + ", "
		}
		verse += "and "
	}

	return verse + gifts[0] + " in a Pear Tree."
}
