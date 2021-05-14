package ctci

import "testing"

func TestFirstCommonAncestor(t *testing.T) {
	root := NewBinaryTreeNode(10, nil, nil, nil)
	root.MassInsert(5, 1, 7, 6, 8, 12)

	testCases := []struct {
		desc         string
		root         *BinaryTreeNode
		firstTarget  int
		secondTarget int
		want         *BinaryTreeNode
	}{
		{
			desc:         "root node is nil",
			root:         nil,
			firstTarget:  1,
			secondTarget: 8,
			want:         nil,
		},
		{
			desc:         "first target node not in tree",
			root:         root,
			firstTarget:  100,
			secondTarget: 8,
			want:         nil,
		},
		{
			desc:         "second target node not in tree",
			root:         root,
			firstTarget:  1,
			secondTarget: 100,
			want:         nil,
		},
		{
			desc:         "targets are on different branches",
			root:         root,
			firstTarget:  1,
			secondTarget: 8,
			want:         root.left,
		},
		{
			desc:         "first target is common ancestor of second",
			root:         root,
			firstTarget:  5,
			secondTarget: 6,
			want:         root.left, // 5
		},
		{
			desc:         "second target is common ancestor of first",
			root:         root,
			firstTarget:  6,
			secondTarget: 5,
			want:         root.left, // 5
		},
		{
			desc:         "target nodes are the same",
			root:         root,
			firstTarget:  12,
			secondTarget: 12,
			want:         root.right,
		},
		{
			desc:         "common ancestor is root",
			root:         root,
			firstTarget:  5,
			secondTarget: 12,
			want:         root,
		},
	}

	for _, tc := range testCases {
		got := tc.root.FirstCommonAncestor(tc.firstTarget, tc.secondTarget)
		if got != tc.want {
			t.Errorf("want common ancestor \n%v, got %v", tc.want, got)
		}

		tc.root.Each(reset)
	}
}

func reset(n *BinaryTreeNode) {
	n.visited = false
}
