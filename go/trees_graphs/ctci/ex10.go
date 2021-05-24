package ctci

import (
	"strconv"
	"strings"
)

// 4.10 Check Substree: T1 and T2 are two very large binary trees, with T1 much
// bigger than T2. Create an algorithm to determine if T2 is a subtree of T1.
//
// A tree T2 is a subtree of T1 if there exists a node n in T1 such that the
// subtree of n is identical to T2. That is, if you cut off the tree at node n,
// the two trees would be identical.
//
// Assumptions:
//   * The trees are not necessarily binary search trees.
//   * The trees may contain duplicate nodes.
//   * The nodes do not know the size of their subtrees.

// CheckSubtree recurses to the bottom of t1 and makes the size of each subtree
// available to the parent node by returning it up the stack. If a node matches
// t2's root and size, the subtree is checked.
//
// Time complexity is O(n + km), where n is the size of t1, m is the size of t2,
// and k is the number of occurences of a node matching the root and size of t2
// in t1.
//
// Space complexity it O(logn + logm), accounting for the recursion to the
// bottom of t1 and any simultaneous comparisons with t2.
func CheckSubtree(t1, t2 *BinaryTreeNode) bool {
	if t1 == nil || t2 == nil {
		return false
	}

	t2Size := t2.Size()

	var checkSubtree func(n *BinaryTreeNode) (int, bool)
	checkSubtree = func(n *BinaryTreeNode) (int, bool) {
		if n == nil {
			return 0, false
		}

		var hasSubtree bool
		var leftSize, rightSize int
		leftSize, hasSubtree = checkSubtree(n.left)
		if hasSubtree {
			return 0, hasSubtree
		}
		rightSize, hasSubtree = checkSubtree(n.right)
		if hasSubtree {
			return 0, hasSubtree
		}

		nSize := leftSize + rightSize + 1
		if nSize == t2Size && n.data == t2.data {
			hasSubtree = match(n, t2)
		}

		return nSize, hasSubtree
	}

	_, isSubtree := checkSubtree(t1)
	return isSubtree
}

func match(n1, n2 *BinaryTreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.data == n2.data {
		return match(n1.left, n2.left) && match(n1.right, n2.right)
	}

	return false
}

// CheckSubtreeString first converts each tree into a string by pre-order
// traversal, which preserves the tree's stucture. It then checks for the
// presence of the substring t2 in t1.
//
// Time complexity is O(n + m), where n is the size of t1 and m is the size of
// t2. This may or may not be faster than CheckSubtree depending on where t2 is
// located within t1, if present, and how many nodes matching t2's root and size
// there are in t1.
//
// Space complexity is O(n + m) to account for the two strings created from the
// trees. This may be significant is the trees have millions of nodes.
func CheckSubtreeString(t1, t2 *BinaryTreeNode) bool {
	var sb1, sb2 strings.Builder
	getOrderString(t1, &sb1)
	getOrderString(t2, &sb2)

	return strings.Contains(sb1.String(), sb2.String())
}

// getOrderString creates a string representation of a binary tree by pre-order
// traversal, which reflects the structure of the input tree. Nil nodes are
// represented with "X", which prevents half-leaf nodes which have a given value
// as either the left or the right child from producing identical strings.
func getOrderString(n *BinaryTreeNode, sb *strings.Builder) {
	if n == nil {
		sb.WriteByte('X')
		return
	}

	sb.WriteString(strconv.Itoa(n.data))
	sb.WriteByte(' ')
	getOrderString(n.left, sb)
	getOrderString(n.right, sb)
}
