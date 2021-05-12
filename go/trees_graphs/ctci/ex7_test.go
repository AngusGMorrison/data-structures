package ctci

import (
	"testing"
)

func TestBuildOrder(t *testing.T) {
	testCases := []struct {
		desc           string
		projects       []string
		deps           []Dependency
		wantBuildOrder []string
		wantErr        error
	}{
		{
			desc:     "dependency graph is acyclic",
			projects: []string{"a", "b", "c", "d", "e", "f"},
			deps: []Dependency{
				{"a", "d"},
				{"f", "b"},
				{"b", "d"},
				{"f", "a"},
				{"d", "c"},
			},
			wantBuildOrder: []string{"f", "a", "b", "d", "c", "e"},
			wantErr:        nil,
		},
		{
			desc:     "dependency graph is cyclic",
			projects: []string{"a", "b"},
			deps: []Dependency{
				{"a", "b"},
				{"b", "a"},
			},
			wantBuildOrder: nil,
			wantErr:        ErrNoBuildOrder,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotBuildOrder, gotErr := BuildOrder(tc.projects, tc.deps)

			if gotErr != tc.wantErr {
				t.Fatalf("want err %v, got %v", tc.wantErr, gotErr)
			}

			if tc.wantBuildOrder == nil {
				if gotBuildOrder != nil {
					t.Errorf("want nil build order, got %+v", gotBuildOrder)
				}
			} else {
				for i, proj := range tc.wantBuildOrder {
					if gotBuildOrder[i] != proj {
						t.Errorf("want build order %+v, got %+v", tc.wantBuildOrder, gotBuildOrder)
					}
				}
			}
		})
	}
}
