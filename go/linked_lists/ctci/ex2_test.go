package ctci

import (
	"fmt"
	"testing"
)

func TestKthToLastIterative(t *testing.T) {
	list := nNodeSinglyLinkedList(5)

	impls := []struct {
		name string
		fn   func(k uint) (int, error)
	}{
		{"KthToLastIterative", list.KthToLastIterative},
		{"KthToLastRecursive", list.KthToLastRecursive},
	}

	testCases := []struct {
		k        uint
		wantData int
		wantErr  error
	}{
		{0, 0, ErrKOutOfBounds},
		{6, 0, ErrKOutOfBounds},
		{1, 4, nil},
		{5, 0, nil},
		{3, 2, nil},
	}

	for _, impl := range impls {
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("%s(%d)", impl.name, tc.k), func(t *testing.T) {
				gotData, gotErr := impl.fn(tc.k)
				if gotErr != tc.wantErr {
					t.Errorf("got err %v, want err %v", gotErr, tc.wantErr)
				}
				if gotData != tc.wantData {
					t.Errorf("got data %v, want data %v", gotData, tc.wantData)
				}

			})
		}
	}
}
