package ctci

// 4.8 First Common Ancestor: Design an algorithm and write code to find the
// first common ancestor of two nodes in a binary tree. Avoid storing additional
// nodes in a data structure. NOTE: This is not necessarily a binary search
// tree.

// FirstCommonAncestor first checks that the first target node exists, and marks
// all nodes along the route to the target node as visited. It then searches for
// the second target node. Once found, it returns up the tree, checking if
// parent nodes are marked 'visited'. The first 'visited' node is the first
// common ancestor.
//
// Requires O(n) time. Requires O(log n) space on the stack for recursive
// calls.
//
// Assumptions
//   * The root of the tree and the data of the target nodes is provided, but
//     the target nodes must be found within the tree.
//   * Nodes do not have knowledge of their parents.
//   * We may record that a node has been visited using a field on the node
//     itself.
func (n *BinaryTreeNode) FirstCommonAncestor(first, second *BinaryTreeNode) *BinaryTreeNode {
	if !n.findAndMarkPath(first) {
		// First node not present in tree.
		return nil
	}

	return n.firstCommonAncestor(second)
}

func (n *BinaryTreeNode) findAndMarkPath(target *BinaryTreeNode) bool {
	if n == nil {
		return false
	}

	// Node found; mark it as visited in case it is the common ancestor.
	if n == target {
		n.visited = true
		return true
	}

	// If either of the node's child branches contains the target node, mark it
	// as seen, tracing the path to the target node through the tree.
	if n.left.findAndMarkPath(target) || n.right.findAndMarkPath(target) {
		n.visited = true
		return true
	}

	return false
}

func (n *BinaryTreeNode) firstCommonAncestor(target *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return nil
	}

	if n == target {
		return n
	}

	for _, child := range []*BinaryTreeNode{n.left, n.right} {
		ancestor := child.firstCommonAncestor(target)
		if ancestor != nil {
			if ancestor.visited {
				return ancestor
			} else {
				return n
			}
		}
	}

	return nil
}

// FirstCommonAncestorCheckBranch assumes that nodes don't know about their
// parents and cannot be marked as visited. It recurses down the tree, checking
// for the point where the target nodes are no longer on the same side of the
// ancestor node. This must be the first common ancestor.
//
// Runtime is O(n). covers is initially called twice on n nodes then, as the
// algorithm branches, 2n/2, 2n/4, etc. times. The constant factor can be optimized (see below).
//
// Space complexity is O(log n).
func (n *BinaryTreeNode) FirstCommonAncestorCheckBranch(first, second *BinaryTreeNode) *BinaryTreeNode {
	// Check that both nodes are present.
	if !n.covers(first) || !n.covers(second) {
		return nil
	}

	return n.findCommonAncestorCheckBranch(first, second)
}

func (n *BinaryTreeNode) findCommonAncestorCheckBranch(first, second *BinaryTreeNode) *BinaryTreeNode {
	if n == nil || n == first || n == second {
		return nil
	}

	firstIsOnLeft := n.left.covers(first)
	secondIsOnLeft := n.left.covers(second)
	if firstIsOnLeft != secondIsOnLeft {
		// Common ancestor found.
		return n
	}

	if firstIsOnLeft {
		return n.left.findCommonAncestorCheckBranch(first, second)
	}
	return n.right.findCommonAncestorCheckBranch(first, second)
}

// FirstCommonAncestorOptimized checks each branch only once, by
// recursing down the tree until the target nodes are located, then returning
// those nodes. When the left and right returned nodes are the first and second
// target nodes, the common ancestor has been found and is returned up the
// stack.
//
// Required O(n) time as the whole tree may have to be searched to find the
// target nodes, but it is faster than FirstCommonAncestorCheckBranch by a
// constant factor.
func (n *BinaryTreeNode) FirstCommonAncestorOptimized(first, second *BinaryTreeNode) *BinaryTreeNode {
	result := n.firstCommonAncestorOptimized(first, second)
	if !result.isCommonAncestor {
		return nil
	}
	return result.node
}

// result contains an isCommonAncestor flag to distinguish cases where one
// target node is not found in the tree. Without it, the other target node would
// be returned and treated as the common ancestor.
type result struct {
	isCommonAncestor bool
	node             *BinaryTreeNode
}

func (n *BinaryTreeNode) firstCommonAncestorOptimized(first, second *BinaryTreeNode) result {
	if n == nil {
		return result{isCommonAncestor: false}
	}

	if n == first && n == second {
		return result{isCommonAncestor: true, node: n}
	}

	leftResult := n.left.firstCommonAncestorOptimized(first, second)
	if leftResult.isCommonAncestor {
		return leftResult
	}

	rightResult := n.right.firstCommonAncestorOptimized(first, second)
	if rightResult.isCommonAncestor {
		return rightResult
	}

	if leftResult.node != nil && rightResult.node != nil {
		// n is the common ancestor
		return result{isCommonAncestor: true, node: n}
	}

	if n == first || n == second {
		// Target node found - check whether it is the ancestor of the other
		// before passing it up the chain.
		isAncestor := leftResult.node != nil || rightResult.node != nil
		return result{isCommonAncestor: isAncestor, node: n}
	}

	// * The common ancestor has not been returned yet;
	// * n is not the common ancestor;
	// * n is not a target node;
	// * Return either the one target node which has been located in the
	//   substree, or nil if no target node has been located.
	res := result{isCommonAncestor: false}
	if leftResult.node != nil {
		res.node = leftResult.node
	} else {
		res.node = rightResult.node
	}

	return res
}

// FirstCommonAncestorWithParents assumes that each node has a link to its
// parent node, and we are given the target nodes as input. It first checks that
// the target nodes are present in the tree. If either the first or second
// covers the other node, then it is the common ancestor and is returned.
// Otherwise, it loops up the tree, checking each of the first node's sibling
// branches for the presence of the second target. The parent of the sibling
// branch must be the first common ancestor.
//
// Takes O(t) time, where t is the number of nodes in the subtree containing the
// first common ancestor. Why not O(t + d), where d is the depth of the first
// common ancestor? Because depth is the logarithm of the number of nodes
// in the tree, meaning the nubmer of nodes in the first common ancestor's
// subtree is the dominant factor. In the worst case, when every node must be
// traversed, runtime in O(n).
//
// Space complexity is O(p), where p is the depth of the deepest target node
// which must be recursively located.
func (n *BinaryTreeNode) FirstCommonAncestorWithParents(first, second *BinaryTreeNode) *BinaryTreeNode {
	// Check if either node is missing from the tree.
	if !n.covers(first) || !n.covers(second) {
		return nil
	}

	// Check if one target node is the common ancestor of both nodes.
	if first.covers(second) {
		return first
	}

	if second.covers(first) {
		return second
	}

	// Traverse upwards from the first node until you find a node that covers the second.
	sibling := first.Sibling()
	parent := first.parent
	for !sibling.covers(second) {
		parent = parent.parent
	}

	return parent
}

// Covers checks for the inclusion of target in binary tree n.
func (n *BinaryTreeNode) covers(target *BinaryTreeNode) bool {
	if n == nil {
		return false
	}

	if n == target {
		return true
	}

	return n.left.covers(target) || n.right.covers(target)
}
