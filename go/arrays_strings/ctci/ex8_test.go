package ctci

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {
	impls := []struct {
		name string
		fn   func([][]int) error
	}{
		{"ZeroWithStorage", ZeroWithStorage},
		{"ZeroNoStorage", ZeroNoStorage},
	}

	testCases := []struct {
		input, want [][]int
		err         error
	}{
		{
			input: [][]int{},
			want:  [][]int{},
			err:   nil,
		},
		{
			input: [][]int{{0}},
			want:  [][]int{{0}},
			err:   nil,
		},
		{
			input: [][]int{{1}},
			want:  [][]int{{1}},
			err:   nil,
		},
		{
			input: [][]int{{1, 0}},
			want:  [][]int{{0, 0}},
			err:   nil,
		},
		{
			input: [][]int{{0, 1}},
			want:  [][]int{{0, 0}},
			err:   nil,
		},
		{
			input: [][]int{{0}, {1}, {1}},
			want:  [][]int{{0}, {0}, {0}},
			err:   nil,
		},
		{
			input: [][]int{{1}, {1}, {1}},
			want:  [][]int{{1}, {1}, {1}},
			err:   nil,
		},
		{
			input: [][]int{{1}, {0}, {1}},
			want:  [][]int{{0}, {0}, {0}},
			err:   nil,
		},
		{
			input: [][]int{
				{1, 2, 0},
				{2, 0},
				{1, 1, 1},
				{2, 2, 2},
			},
			want: [][]int{
				{1, 2, 0},
				{2, 0},
				{1, 1, 1},
				{2, 2, 2},
			},
			err: ErrRowLenMismatch,
		},
		{
			input: [][]int{
				{1, 2, 0},
				{2, 0, 1},
				{1, 1, 1},
				{2, 2, 2},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{1, 0, 0},
				{2, 0, 0},
			},
			err: nil,
		},
		{
			input: [][]int{
				{1, 2, 0},
				{2, 0, 1},
				{0, 1, 1},
				{2, 2, 2},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			err: nil,
		},
		{
			input: [][]int{
				{1, 2, 0},
				{2, 2, 1},
				{0, 1, 1},
				{2, 2, 2},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 2, 0},
				{0, 0, 0},
				{0, 2, 0},
			},
			err: nil,
		},
	}

	for _, impl := range impls {
		for _, tc := range testCases {
			input := clone(tc.input)
			t.Run(fmt.Sprintf("%s(%+v)", impl.name, input), func(t *testing.T) {
				err := impl.fn(input)
				if err != tc.err {
					t.Errorf("want err %v, got %v", tc.err, err)
				}

				if !reflect.DeepEqual(input, tc.want) {
					t.Errorf("want matrix %+v, got %+v", tc.want, input)
				}
			})
		}
	}
}

func clone(matrix [][]int) [][]int {
	cp := make([][]int, len(matrix))
	for i, row := range matrix {
		cp[i] = make([]int, len(row))
		copy(cp[i], row)
	}

	return cp
}
