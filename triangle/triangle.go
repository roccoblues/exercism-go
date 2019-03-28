package triangle

import "math"

// Kind is the triangle type
type Kind int

const (
	// NaT means not a triangle
	NaT = iota
	// Equ is an equilateral triangle
	Equ
	// Iso is an isosceles triangle
	Iso
	// Sca is an scalene triangle
	Sca
)

// KindFromSides determines the type of a triangle.
func KindFromSides(a, b, c float64) Kind {
	if inValidSide(a) || inValidSide(b) || inValidSide(c) {
		return NaT
	}

	// the sum of the lengths of any two sides must be greater than or equal to
	// the length of the third side
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// equilateral triangle has all three sides the same length
	if a == b && b == c {
		return Equ
	}

	// isosceles triangle has at least two sides the same length
	if a == b || a == c || b == c {
		return Iso
	}

	// scalene triangle has all sides of different lengths
	return Sca
}

func inValidSide(s float64) bool {
	return math.IsNaN(s) || math.IsInf(s, 0) || s <= 0
}
