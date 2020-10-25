package ctci

import (
	"fmt"
	"testing"
)

func TestIsRotation(t *testing.T) {
	testCases := []struct {
		s1, s2 string
		want   bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbottle", "tlewaterbot", true},
		{"", "", true},
		{"abc", "ab", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("IsRotation(%q, %q)", tc.s1, tc.s2), func(t *testing.T) {
			got := IsRotation(tc.s1, tc.s2)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
