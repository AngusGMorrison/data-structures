package rabinkarp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllIndexes(t *testing.T) {
	testCases := []struct {
		pat, text string
		wantErr   error
		wantIdxs  []int
	}{
		{"", "doearehearingme", ErrNoPattern, nil},
		{"doearehearingmedoearehearingme", "doearehearingme", ErrPatExceedsTextLen, nil},
		{"doearehearingme", "doearehearingme", nil, []int{0}},
		{"ear", "doearehearingme", nil, []int{2, 7}},
		{"end", "attheend", nil, []int{5}},
		{"a", "aaaaa", nil, []int{0, 1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("AllIndexes(%q, %q)", tc.pat, tc.text), func(t *testing.T) {
			idxs, err := AllIndexes(tc.pat, tc.text)
			if err != tc.wantErr {
				t.Errorf("got err %v, want err %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(idxs, tc.wantIdxs) {
				t.Errorf("got idxs %+v, want idxs %+v", idxs, tc.wantIdxs)
			}

		})
	}
}
