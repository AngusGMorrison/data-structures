package ctci

import "testing"

func TestInOrderSuccessorLoop(t *testing.T) {
	root := NewBinaryTreeNode(50, nil, nil, nil)
	root.MassInsert(40, 30, 45, 42, 48, 60, 70)

	testCases := []struct {
		desc        string
		input, want *BinaryTreeNode
	}{
		{
			desc:  "node is nil",
			input: nil,
			want:  nil,
		},
		{
			desc:  "node is greatest in tree",
			input: root.right.right,
			want:  nil,
		},
		{
			desc:  "in-order successor is right child",
			input: root,
			want:  root.right,
		},
		{
			desc:  "in-order successor is leftmost descendant of right child",
			input: root.left,
			want:  root.left.right.left,
		},
		{
			desc:  "in-order successor is parent",
			input: root.left.left,
			want:  root.left,
		},
		{
			desc:  "in-order successor is distant ancestor",
			input: root.left.right.right,
			want:  root,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.input.InOrderSuccessorLoop()
			if got != tc.want {
				t.Errorf("want %d, got %d", tc.want.data, got.data)
			}
		})
	}
}

func TestInOrderSuccessorRecursive(t *testing.T) {
	root := NewBinaryTreeNode(50, nil, nil, nil)
	root.MassInsert(40, 30, 45, 42, 48, 60, 70)

	testCases := []struct {
		desc        string
		input, want *BinaryTreeNode
	}{
		{
			desc:  "node is nil",
			input: nil,
			want:  nil,
		},
		{
			desc:  "node is greatest in tree",
			input: root.right.right,
			want:  nil,
		},
		{
			desc:  "in-order successor is right child",
			input: root,
			want:  root.right,
		},
		{
			desc:  "in-order successor is leftmost descendant of right child",
			input: root.left,
			want:  root.left.right.left,
		},
		{
			desc:  "in-order successor is parent",
			input: root.left.left,
			want:  root.left,
		},
		{
			desc:  "in-order successor is distant ancestor",
			input: root.left.right.right,
			want:  root,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.input.InOrderSuccessorRecursive()
			if got != tc.want {
				t.Errorf("want %d, got %d", tc.want.data, got.data)
			}
		})
	}
}
