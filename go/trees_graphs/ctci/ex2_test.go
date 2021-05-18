package ctci

import (
	"testing"
)

func TestTreeFromSlice(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		want  string
	}{
		{
			desc:  "slice has even number of elements",
			input: []int{10, 20, 30, 40, 50, 60, 70, 80},
			want:  "\t\t80\n\t70\n\t\t60\n50\n\t\t40\n\t30\n\t\t20\n\t\t\t10\n",
		},
		{
			desc:  "slice has odd number of elements",
			input: []int{20, 30, 40, 50, 60, 70, 80},
			want:  "\t\t80\n\t70\n\t\t60\n50\n\t\t40\n\t30\n\t\t20\n",
		},
		{
			desc:  "slice is empty",
			input: []int{},
			want:  "",
		},
		{
			desc:  "slice is nil",
			input: nil,
			want:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := TreeFromSlice(tc.input).PrintTree()
			if got != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
			}

		})
	}
}
