package ctci

import (
	"errors"
)

// 1.7 Rotate Matrix: Given an image represented by an NxN matrix, where each
// pixel in the image is 4 bytes, write a method to rotate the image by 90
// degrees. Can you do this in place?
//
// Note: Implemented using integers rather than byte arrays for ease of testing.
//
// Assumptions:
//	* The direction of rotation.

// RotateClockwise rotates a 2D matrix clockwise by 90 degrees, treating the
// matrix as a series on concentric layers or "rings". Each layer is rotated
// in turn, starting from the outermost ring and moving inwards.
//
// Time complexity: O(n^2); every element must be touched
// Space complexity: O(1)
func RotateClockwise(matrix [][]int) error {
	size := len(matrix)
	if size == 0 {
		return ErrZeroLenMatrix
	}

	for _, subarr := range matrix {
		if len(subarr) != size {
			return ErrSubarrayLenMismatch
		}
	}

	// Rotate the matrix in layers, from the outermost "ring" of elements to
	// the innermost.
	for layer := 0; layer < size/2; layer++ {
		first := layer           // the top and left-hand bound of the current layer
		last := size - 1 - layer // the bottom and right-hand bound of the current layer

		// Depending on whether it is used as the y- or x-coordinate when
		// accessing elements, i represents increasing vertical or horizontal
		// progress repsectively.
		for i := first; i < last; i++ {
			// Offset is the equivalent of i when "progress" is measured
			// relative to the end of the column or row
			offset := i - first

			// Save the element to be rotated from the top row.
			top := matrix[first][i]
			// Left -> top.
			matrix[first][i] = matrix[last-offset][first]
			// Bottom -> left.
			matrix[last-offset][first] = matrix[last][last-offset]
			// Right -> bottom.
			matrix[last][last-offset] = matrix[i][last]
			// Top -> right.
			matrix[i][last] = top
		}
	}

	return nil
}

// RotateAnticlockwise uses the same algorithm as RotateClockwise, but rotates
// in the opposite direction.
func RotateAnticlockwise(matrix [][]int) error {
	size := len(matrix)
	if size == 0 {
		return ErrZeroLenMatrix
	}

	for _, subarr := range matrix {
		if len(subarr) != size {
			return ErrSubarrayLenMismatch
		}
	}

	for layer := 0; layer < size/2; layer++ {
		first := layer
		last := size - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first

			// Save top entry.
			top := matrix[first][last-offset]
			// Right -> top.
			matrix[first][last-offset] = matrix[last-offset][last]
			// Bottom -> right.
			matrix[last-offset][last] = matrix[last][i]
			// Left -> bottom.
			matrix[last][i] = matrix[i][first]
			// Top - left.
			matrix[i][first] = top
		}
	}
	return nil
}

var ErrZeroLenMatrix = errors.New("Matrix length cannot be zero")
var ErrSubarrayLenMismatch = errors.New("all subarrays must have the same length")
