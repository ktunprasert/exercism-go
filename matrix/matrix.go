package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	strRows := strings.Split(s, "\n")
	m := make(Matrix, len(strRows))
	var cols int
	for i, row := range strRows {
		f := strings.Fields(row)
		if i == 0 {
			cols = len(f)
		}

		if cols != len(f) {
			return nil, errors.New("invalid matrix")
		}

		r := make([]int, len(f))

		for j, v := range f {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			r[j] = num

		}

		m[i] = r
	}

	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	matrix := make([][]int, len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		row := make([]int, 0)

		for j := 0; j < len(m); j++ {
			row = append(row, m[j][i])
		}

		matrix[i] = row
	}

	return matrix
}

func (m Matrix) Rows() [][]int {
	matrix := make([][]int, len(m))

	for i, rows := range m {
		matrix[i] = make([]int, len(rows))
		copy(matrix[i], rows)
	}

	return matrix
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}

	m[row][col] = val
	return true
}
