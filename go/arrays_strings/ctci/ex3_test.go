package ctci

import (
	"fmt"
	"testing"
)

func TestURLify(t *testing.T) {
	testCases := []struct {
		input  string
		spaces int
		want   string
	}{
		{"", 0, ""},
		{"Mr John Smith", 2, "Mr%20John%20Smith"},
		{" Mr John Smith ", 4, "%20Mr%20John%20Smith%20"},
		{"\fMr\tJohn\nSmith\r\v", 0, "\fMr\tJohn\nSmith\r\v"},
	}

	for _, tc := range testCases {
		length := len(tc.input)
		t.Run(fmt.Sprintf("URLify(%s, %d)", tc.input, length), func(t *testing.T) {
			capacity := length + tc.spaces*2
			runes := make([]rune, capacity)
			for i, r := range []rune(tc.input) {
				runes[i] = r
			}

			got := string(URLify(runes, length))
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}

		})
	}
}
