// Package clock implements a clock that handles times without dates.
package clock

import (
	"fmt"
)

// Clock represents a clock with hours and minutes.
type Clock struct {
	hour   int
	minute int
}

// New returns a new Clock.
func New(hour, minute int) Clock {
	for minute < 0 {
		minute = 60 + minute
		hour--
	}
	for minute >= 60 {
		minute -= 60
		hour++
	}

	for hour < 0 {
		hour = 24 + hour
	}
	for hour >= 24 {
		hour -= 24
	}

	return Clock{hour, minute}
}

// String returns the string representation of the clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add returns a new Clock with the given minutes added.
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract returns a new Clock with the given minutes subtracted.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
