package ctci

import (
	"fmt"
	"testing"
)

func TestAllUniqueChars(t *testing.T) {
	impls := []struct {
		name string
		fn   func(string) bool
	}{
		{"AllUniqueCharsBruteForce", AllUniqueCharsBruteForce},
		{"AllUniqueCharsArr", AllUniqueCharsArr},
		{"AllUniqueCharsBV", AllUniqueCharsBV},
		{"AllUniqueCharsSort", AllUniqueCharsSort},
	}

	for _, impl := range impls {
		testCases := []struct {
			input string
			want  bool
		}{
			{"", true},
			{"a", true},
			{"abc", true},
			{"abac", false},
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s(%q)", impl.name, tc.input), func(t *testing.T) {
				if got := impl.fn(tc.input); got != tc.want {
					t.Errorf("got %t, want %t", got, tc.want)
				}
			})
		}
	}
}
