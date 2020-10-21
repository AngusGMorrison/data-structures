package ctci

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	impls := []struct {
		name string
		fn   func(string) string
	}{
		{"CompressOnePass", CompressOnePass},
		{"CompressTwoPass", CompressTwoPass},
	}

	testCases := []struct {
		input, want string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"aabbccdd", "aabbccdd"},
		{"", ""},
		{"a", "a"},
		{"ZZZzzz", "Z3z3"},
	}

	for _, impl := range impls {
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s(%q)", impl.name, tc.input), func(t *testing.T) {
				got := impl.fn(tc.input)
				if got != tc.want {
					t.Errorf("got %q, want %q", got, tc.want)
				}
			})
		}
	}
}
