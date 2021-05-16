package ctci

import "testing"

type testCase struct {
	desc         string
	root         *BinaryTreeNode
	firstTarget  *BinaryTreeNode
	secondTarget *BinaryTreeNode
	want         *BinaryTreeNode
}

func TestFirstCommonAncestor(t *testing.T) {
	for _, tc := range testCases() {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.root.FirstCommonAncestorWithParents(tc.firstTarget, tc.secondTarget)
			if got != tc.want {
				t.Errorf("want common ancestor \n%v, got %v", tc.want, got)
			}

			tc.root.Each(reset)
		})
	}
}

func TestFirstCommonAncestorWithParents(t *testing.T) {
	for _, tc := range testCases() {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.root.FirstCommonAncestorWithParents(tc.firstTarget, tc.secondTarget)
			if got != tc.want {
				t.Errorf("want common ancestor \n%v, got %v", tc.want, got)
			}

			tc.root.Each(reset)
		})
	}
}

func TestFirstCommonAncestorCheckBranch(t *testing.T) {
	for _, tc := range testCases() {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.root.FirstCommonAncestorWithParents(tc.firstTarget, tc.secondTarget)
			if got != tc.want {
				t.Errorf("want common ancestor \n%v, got %v", tc.want, got)
			}

			tc.root.Each(reset)
		})
	}
}

func TestFirstCommonAncestorOptimized(t *testing.T) {
	for _, tc := range testCases() {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.root.FirstCommonAncestorOptimized(tc.firstTarget, tc.secondTarget)
			if got != tc.want {
				t.Errorf("want common ancestor \n%v, got %v", tc.want, got)
			}

			tc.root.Each(reset)
		})
	}
}

func testCases() []testCase {
	root := NewBinaryTreeNode(10, nil, nil, nil)
	root.MassInsert(5, 1, 7, 6, 8, 12)

	return []testCase{
		{
			desc:         "root node is nil",
			root:         nil,
			firstTarget:  root.left.left,
			secondTarget: root.left.right.right,
			want:         nil,
		},
		{
			desc:         "first target node not in tree",
			root:         root,
			firstTarget:  NewBinaryTreeNode(100, nil, nil, nil),
			secondTarget: root.left.right.right,
			want:         nil,
		},
		{
			desc:         "second target node not in tree",
			root:         root,
			firstTarget:  root.left.left,
			secondTarget: NewBinaryTreeNode(100, nil, nil, nil),
			want:         nil,
		},
		{
			desc:         "targets are on different branches",
			root:         root,
			firstTarget:  root.left.left,
			secondTarget: root.left.right.right,
			want:         root.left,
		},
		{
			desc:         "first target is common ancestor of second",
			root:         root,
			firstTarget:  root.left,
			secondTarget: root.left.right.left,
			want:         root.left, // 5
		},
		{
			desc:         "second target is common ancestor of first",
			root:         root,
			firstTarget:  root.left.right.left,
			secondTarget: root.left,
			want:         root.left, // 5
		},
		{
			desc:         "target nodes are the same",
			root:         root,
			firstTarget:  root.right,
			secondTarget: root.right,
			want:         root.right,
		},
		{
			desc:         "common ancestor is root",
			root:         root,
			firstTarget:  root.left,
			secondTarget: root.right,
			want:         root,
		},
	}
}

func reset(n *BinaryTreeNode) {
	n.visited = false
}
