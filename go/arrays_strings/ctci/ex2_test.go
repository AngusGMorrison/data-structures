package ctci

import (
	"fmt"
	"testing"
)

func TestIsPermutationArr(t *testing.T) {
	impls := []struct {
		name string
		fn   func(s1, s2 string) bool
	}{
		{"IsPermutationArr", IsPermutationArr},
		{"IsPermutationSort", IsPermutationSort},
	}

	for _, impl := range impls {
		testCases := []struct {
			s1, s2 string
			want   bool
		}{
			{"", "abc", false},
			{"sdasdasdasert", "abcfb", false},
			{"TEST", "test", false},
			{"t e s t", "test", false},
			{"", "", true},
			{"example string", "genxiarmptlse ", true},
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s(%q, %q)", impl.name, tc.s1, tc.s2), func(t *testing.T) {
				got := impl.fn(tc.s1, tc.s2)
				if got != tc.want {
					t.Errorf("got %t, want %t", got, tc.want)
				}
			})
		}
	}
}
