package ctci

import (
	"fmt"
	"testing"
)

func TestIsPalindromePerm(t *testing.T) {
	impls := []struct {
		name string
		fn   func(string) bool
	}{
		{"IsPalindromePerm", IsPalindromePerm},
		{"IsPalindromePermAlt", IsPalindromePermAlt},
		{"IsPalindromePermVector", IsPalindromePermVector},
	}

	testCases := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"tact coa", true},
		{"Tact coa", true},
		{"Tac1tc1oa", true},
		{"\tTact\ncoa\f", true},
		{"Tact ecoa", false},
		{"abcdef", false},
	}

	for _, impl := range impls {
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s(%q)", impl.name, tc.s), func(t *testing.T) {
				got := impl.fn(tc.s)
				if got != tc.want {
					t.Errorf("got %t, want %t", got, tc.want)
				}
			})
		}
	}
}
