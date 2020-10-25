package ctci

import "errors"

// 1.8 Zero Matrix: Write an algorithm such that if an element in an MxN matrix
// is 0, its entire row and column are set to 0.

// ZeroWithStorage uses auxiliary boolean arrays to flag which rows and columns
// should be zeroed. The first pass over the matrix populates the arrays,
// and the second sets the relevant elements in the matrix to zero. This can't
// be done in a single pass because whole columns to zero early in the iteration
// will cause each subsequent row to be set to zero, clearing most of the array.
//
// Time complexity: O(mn)
// Space complexity: O(m+n)
func ZeroWithStorage(matrix [][]int) error {
	if len(matrix) == 0 {
		return nil
	}

	rowLen := len(matrix[0])
	rows := make([]bool, len(matrix))
	cols := make([]bool, rowLen)

	for i, row := range matrix {
		if len(row) != rowLen {
			return ErrRowLenMismatch
		}

		for j, elem := range row {
			if elem == 0 {
				rows[i], cols[j] = true, true
			}
		}
	}

	for i, zero := range rows {
		if zero {
			nullifyRow(matrix, i)
		}
	}

	for i, zero := range cols {
		if zero {
			nullifyCol(matrix, i)
		}
	}

	return nil
}

// ZeroNoStorage uses the first row and column of the input matrix to flag
// whether the rest of that row and column should be zeroed.
//
// Time complexity: O(mn)
// Space complexity: O(1)
func ZeroNoStorage(matrix [][]int) error {
	if len(matrix) == 0 {
		return nil
	}

	rowLen := len(matrix[0])
	// Record whether the first row and column start with zeroes in them. As
	// the first row and column are used as a flag store indicating whether
	// other rows and columns should be zeroed, they must zeroed last, after
	// every other row and column has been updated.
	var zeroFirstCol, zeroFirstRow bool
	for _, row := range matrix {
		if len(row) != rowLen {
			return ErrRowLenMismatch
		}

		if row[0] == 0 {
			zeroFirstCol = true
			break
		}
	}

	for _, elem := range matrix[0] {
		if elem == 0 {
			zeroFirstRow = true
			break
		}
	}

	// For every row and column apart from the first, check every element for
	// 0 values. If 0, set the corresponding cell in the first row and column
	// to zero, indicating that all values in this row and column should be
	// zeroed.
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < rowLen; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	// For every row except the first, check whether the first column is 0. If
	// so zero out the whole row.
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			nullifyRow(matrix, i)
		}
	}

	// For every column except the first, check whether the first row is 0. If
	// so, zero out the whole column.
	for i := 1; i < rowLen; i++ {
		if matrix[0][i] == 0 {
			nullifyCol(matrix, i)
		}
	}

	// Set first row to zero.
	if zeroFirstRow {
		nullifyRow(matrix, 0)
	}

	// Set first column to zero.
	if zeroFirstCol {
		nullifyCol(matrix, 0)
	}

	return nil
}

func nullifyRow(matrix [][]int, row int) {
	for i := range matrix[row] {
		matrix[row][i] = 0
	}
}

func nullifyCol(matrix [][]int, col int) {
	for i := range matrix {
		matrix[i][col] = 0
	}
}

// ErrRowLenMismatch indicates that the provided matrix contains rows of
// varying length.
var ErrRowLenMismatch = errors.New("all rows must have the same length")
