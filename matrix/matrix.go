// Package matrix implements a two-dimensional matrix.
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix as a two-dimensional array.
type Matrix [][]int

// New returns a new matrix from the given input.
func New(input string) (Matrix, error) {
	numRows := strings.Count(input, "\n") + 1

	firstRow := strings.SplitN(input, "\n", 2)[0]
	numCols := strings.Count(firstRow, " ") + 1

	m := make(Matrix, numRows)
	for i := 0; i < numRows; i++ {
		m[i] = make([]int, numCols)
	}

	for i, row := range strings.Split(input, "\n") {
		for j, val := range strings.Split(strings.Trim(row, " "), " ") {
			if j >= numCols {
				return nil, fmt.Errorf("New: uneven columns")
			}
			v, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			m[i][j] = v
		}
	}

	return m, nil
}

// Rows returns a list of the rows, reading each row left-to-right while
// moving top-to-bottom across the rows.
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))

	for i := 0; i < len(m); i++ {
		rows[i] = make([]int, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}

	return rows
}

// Cols returns a list of the columns, reading each column top-to-bottom while
// moving from left-to-right.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(m); j++ {
			cols[i][j] = m[j][i]
		}
	}

	return cols
}

// Set the given matrix element.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[row]) {
		return false
	}
	m[row][col] = val

	return true
}
