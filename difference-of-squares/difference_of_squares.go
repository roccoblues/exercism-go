// Package diffsquares implements methods to calculate the difference of squares.
package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	// Sum of the First n Natural Numbers can be calculated with: n*(n+1)/2
	sum := n * (n + 1) / 2

	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first N natural numbers.
func SumOfSquares(n int) int {
	// Sum of the Squares of the First n Natural Numbers can be calculated with:  n*(n+1)*(2*n+1)/6
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between the square of the sum and the
// sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
