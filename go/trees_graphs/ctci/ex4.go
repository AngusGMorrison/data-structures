package ctci

// 4.4 Check Balanced: Implement a function to check if a binary tree is
// balanced. For the purposes of this question, a balanced tree is defined to be
// a tree such that the heights of the two subtrees of any node never differ by
// more than one.

// CheckBalanced checks whether a binary tree is balanced by recursing to the
// bottom of the tree and declaring the height of each leaf node to be one. The
// heights of the branches are returned to their parents, which check for
// balance before incrementing their own height.
//
// Requires O(1) space and O(n) time, as every node must be visited in the case
// of a balanced tree.
func CheckBalanced(root *BinaryTreeNode) bool {
	_, balanced := heightAndBalance(root)
	return balanced
}

func heightAndBalance(n *BinaryTreeNode) (int, bool) {
	if n == nil {
		return 0, true
	}

	lHeight, lBalanced := heightAndBalance(n.left)
	rHeight, rBalanced := heightAndBalance(n.right)

	if !lBalanced || !rBalanced {
		return 0, false
	}

	heightDiff := abs(lHeight - rHeight)
	if heightDiff > 1 {
		return 0, false
	}

	nHeight := max(lHeight, rHeight) + 1
	return nHeight, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
