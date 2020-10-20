package ctci

import "testing"

func TestMaxOneAway(t *testing.T) {
	testCases := []struct {
		s1, s2 string
		want   bool
	}{
		{"pale", "pale", true},
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bales", false},
		{"pale", "bake", false},
		{"", "b", true},
		{"", "bc", false},
	}

	for _, tc := range testCases {
		got := MaxOneAway(tc.s1, tc.s2)
		if got != tc.want {
			t.Errorf("got %t, want %t", got, tc.want)
		}
	}
}
