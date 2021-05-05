package ctci

import "math"

// 4.5 Validate BST: Implement a function to check if a binary tree is a binary
// search tree.

// IsBST recusively checks whether each left child is nil or less than smallest
// element seen so far, and whether each right child is nil or greater than the
// largest element seen so far. Assumptions:
//   * The tree may or may not be balanced.
//   * A nil tree is considered a binary search tree.
// Takes O(n) time and O(log n) space to hold recusive calls to IsBST().
func (n *BinaryTreeNode) IsBST() bool {
	return n.isBST(nil, nil)
}

func (n *BinaryTreeNode) isBST(min, max *BinaryTreeNode) bool {
	if n == nil {
		return true
	}

	if min != nil {
		if n.data > min.data {
			return false
		}
	}

	if max != nil {
		if n.data <= max.data {
			return false
		}
	}

	return n.left.isBST(n, max) && n.right.isBST(min, n)
}

// IsBSTInOrder performs an in-order traversal of the input tree, verfiying that
// each element is equal to or less than the previous element in the tree. Also
// takes O(n) time and O(log n) space.
func (n *BinaryTreeNode) IsBSTInOrder() bool {
	_, inOrder := n.isBSTInOrder(int(math.MinInt64))
	return inOrder
}

func (n *BinaryTreeNode) isBSTInOrder(prevData int) (data int, inOrder bool) {
	if n == nil {
		return prevData, true
	}

	prevData, inOrder = n.left.isBSTInOrder(prevData)
	if !inOrder || prevData > n.data {
		return 0, false
	}

	return n.right.isBSTInOrder(n.data)
}
