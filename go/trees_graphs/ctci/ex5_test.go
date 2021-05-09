package ctci

import "testing"

func TestIsBST(t *testing.T) {
	notBST := NewBinaryTreeNode(20, nil, nil, nil)
	notBST.left = NewBinaryTreeNode(10, nil, nil, nil)
	notBST.right = NewBinaryTreeNode(30, nil, nil, nil)
	notBST.left.right = NewBinaryTreeNode(25, nil, nil, nil)

	bst := NewBinaryTreeNode(50, nil, nil, nil)
	bst.MassInsert(40, 60, 55, 10)

	testCases := []struct {
		desc  string
		input *BinaryTreeNode
		want  bool
	}{
		{
			desc:  "nil tree",
			input: nil,
			want:  true,
		},
		{
			desc:  "invalid BST",
			input: notBST,
			want:  false,
		},
		{
			desc:  "valid BST",
			input: bst,
			want:  true,
		},
	}

	for _, tc := range testCases {
		got := tc.input.IsBST()
		if got != tc.want {
			t.Errorf("want %t, got %t", tc.want, got)
		}
	}
}

func TestIsBSTInOrder(t *testing.T) {
	notBST := NewBinaryTreeNode(20, nil, nil, nil)
	notBST.left = NewBinaryTreeNode(10, nil, nil, nil)
	notBST.right = NewBinaryTreeNode(30, nil, nil, nil)
	notBST.left.right = NewBinaryTreeNode(25, nil, nil, nil)

	bst := NewBinaryTreeNode(50, nil, nil, nil)
	bst.MassInsert(40, 60, 55, 10)

	testCases := []struct {
		desc  string
		input *BinaryTreeNode
		want  bool
	}{
		{
			desc:  "nil tree",
			input: nil,
			want:  true,
		},
		{
			desc:  "invalid BST",
			input: notBST,
			want:  false,
		},
		{
			desc:  "valid BST",
			input: bst,
			want:  true,
		},
	}

	for _, tc := range testCases {
		got := tc.input.IsBSTInOrder()
		if got != tc.want {
			t.Errorf("want %t, got %t", tc.want, got)
		}
	}
}
