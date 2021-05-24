package ctci

import "testing"

func TestCheckSubtree(t *testing.T) {
	for _, tc := range subtreeTestCases() {
		t.Run(tc.desc, func(t *testing.T) {
			got := CheckSubtree(tc.t1, tc.t2)
			if got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func TestCheckSubtreeString(t *testing.T) {
	for _, tc := range subtreeTestCases() {
		t.Run(tc.desc, func(t *testing.T) {
			got := CheckSubtreeString(tc.t1, tc.t2)
			if got != tc.want {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

type subtreeTestCase struct {
	desc   string
	t1, t2 *BinaryTreeNode
	want   bool
}

func subtreeTestCases() []subtreeTestCase {
	return []subtreeTestCase{
		{
			desc: "t1 is nil",
			t1:   nil,
			t2:   buildTree([]int{10, 7, 15}),
			want: false,
		},
		{
			desc: "t2 is nil",
			t1:   nil,
			t2:   buildTree([]int{10, 7, 15}),
			want: false,
		},
		{
			desc: "subtree is present and tree contains no duplicates",
			t1:   buildTree([]int{2, 1, 3, 10, 7, 15}),
			t2:   buildTree([]int{10, 7, 15}),
			want: true,
		},
		{
			desc: "subtree is absent and tree contains no duplicates",
			t1:   buildTree([]int{2, 1, 3, 10, 7, 15}),
			t2:   buildTree([]int{10, 7, 11}),
			want: false,
		},
		{
			desc: "subtree is present and tree contains duplicates",
			t1:   buildTree([]int{50, 70, 30, 20, 40, 30, 10, 25}),
			t2:   buildTree([]int{30, 25}),
			want: true,
		},
		{
			desc: "subtree is absent and tree contains duplicates",
			t1:   buildTree([]int{50, 70, 30, 20, 40, 30, 10, 25}),
			t2:   buildTree([]int{30, 22}),
			want: false,
		},
	}
}
