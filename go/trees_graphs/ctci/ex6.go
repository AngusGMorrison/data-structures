package ctci

// 4.6 Successor: Write an algorithm to find the next node (i.e. the in-order
// successor) of a given node in a binary search tree. You may assume that each
// node has a link to its parent.

// InOrderSuccessorLoop finds the next node which is greater than or equal to
// the value of the given node by first checking for the presence of a right
// child and returning its leftmost child. If there is no right child, it loops
// up the tree until it finds a parent for which the input node is a left
// descendant.
//
// Requires O(n/2) time. In the worst-case scenario, when the input node is the
// greatest in the tree, the algorithm must recurse back to the root node before
// returning. Requires only O(1) space using a loop instead of recursion.
//
// Assumptions:
//   * The in-order successor may have the same numeric value as the original
//     node.
func (n *BinaryTreeNode) InOrderSuccessorLoop() *BinaryTreeNode {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return n.right.leftmostChild()
	}

	prev, cur := n, n.parent
	for cur != nil && cur.left != prev {
		prev, cur = cur, cur.parent
	}

	return cur
}

// InOrderSuccessorRecursive finds the next node which is greater than or equal
// to the value of the given node by first checking for the presence of a right
// child and returning its leftmost child. If there is no right child, it
// recurses up the tree until it finds a parent for which the input node is a
// left descendant.
//
// Requires O(n/2) time. In the worst-case scenario, when the input node is the
// greatest in the tree, the algorithm must recurse back to the root node before
// returning. Requires O(log n) space for recursive calls on the stack.
//
// Assumptions:
//   * The in-order successor may have the same numeric value as the original
//     node.
func (n *BinaryTreeNode) InOrderSuccessorRecursive() *BinaryTreeNode {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return n.right.leftmostChild()
	}

	return n.inOrderSuccessor(n)
}

func (n *BinaryTreeNode) inOrderSuccessor(prev *BinaryTreeNode) *BinaryTreeNode {
	// Base case: no in-order successor found.
	if n == nil {
		return nil
	}

	// If the node we've just come from is the left child of the current node,
	// the current node must be the in-order successor.
	if n.left == prev {
		return n
	}

	// Move one level up the tree and check again.
	return n.parent.inOrderSuccessor(n)
}

// leftMostChild returns the leftmost descendant of n, or n itself if it has no
// left branch.
func (n *BinaryTreeNode) leftmostChild() *BinaryTreeNode {
	if n == nil {
		return nil
	}

	var cur *BinaryTreeNode
	for cur = n; cur.left != nil; cur = cur.left {
	}

	return cur
}
