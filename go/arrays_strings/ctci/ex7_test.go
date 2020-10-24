package ctci

import (
	"reflect"
	"testing"
)

func TestRotateClockwise(t *testing.T) {
	testCases := []struct {
		input, want [][]int
		err         error
	}{
		{
			input: [][]int{},
			want:  [][]int{},
			err:   ErrZeroLenMatrix,
		},
		{
			input: [][]int{
				{1, 1, 1, 2},
				{4, 5, 6},
				{4},
				{4, 3, 3, 3},
			},
			want: [][]int{
				{1, 1, 1, 2},
				{4, 5, 6},
				{4},
				{4, 3, 3, 3},
			},
			err: ErrSubarrayLenMismatch,
		},
		{
			input: [][]int{
				{1, 2},
				{4, 3},
			},
			want: [][]int{
				{4, 1},
				{3, 2},
			},
			err: nil,
		},
		{
			input: [][]int{
				{1, 1, 2},
				{4, 5, 2},
				{4, 3, 3},
			},
			want: [][]int{
				{4, 4, 1},
				{3, 5, 1},
				{3, 2, 2},
			},
			err: nil,
		},
		{
			input: [][]int{
				{1, 1, 1, 2},
				{4, 5, 6, 2},
				{4, 8, 7, 2},
				{4, 3, 3, 3},
			},
			want: [][]int{
				{4, 4, 4, 1},
				{3, 8, 5, 1},
				{3, 7, 6, 1},
				{3, 2, 2, 2},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		err := RotateClockwise(tc.input)
		if err != tc.err {
			t.Errorf("want err %v, got %v", tc.err, err)
		}

		if !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("want rotated matrix %+v, got %+v", tc.want, tc.input)
		}
	}
}

func TestRotateAnticlockwise(t *testing.T) {
	testCases := []struct {
		input, want [][]int
		err         error
	}{
		{
			input: [][]int{},
			want:  [][]int{},
			err:   ErrZeroLenMatrix,
		},
		{
			input: [][]int{
				{1, 1, 1, 2},
				{4, 5, 6},
				{4},
				{4, 3, 3, 3},
			},
			want: [][]int{
				{1, 1, 1, 2},
				{4, 5, 6},
				{4},
				{4, 3, 3, 3},
			},
			err: ErrSubarrayLenMismatch,
		},
		{
			input: [][]int{
				{1, 2},
				{4, 3},
			},
			want: [][]int{
				{2, 3},
				{1, 4},
			},
			err: nil,
		},
		{
			input: [][]int{
				{1, 1, 2},
				{4, 5, 2},
				{4, 3, 3},
			},
			want: [][]int{
				{2, 2, 3},
				{1, 5, 3},
				{1, 4, 4},
			},
			err: nil,
		},
		{
			input: [][]int{
				{1, 1, 1, 2},
				{4, 5, 6, 2},
				{4, 8, 7, 2},
				{4, 3, 3, 3},
			},
			want: [][]int{
				{2, 2, 2, 3},
				{1, 6, 7, 3},
				{1, 5, 8, 3},
				{1, 4, 4, 4},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		err := RotateAnticlockwise(tc.input)
		if err != tc.err {
			t.Errorf("want err %v, got %v", tc.err, err)
		}

		if !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("want rotated matrix %+v, got %+v", tc.want, tc.input)
		}
	}
}
