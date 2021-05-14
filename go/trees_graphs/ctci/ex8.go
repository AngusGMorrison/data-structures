package ctci

// 4.8 First Common Ancestor: Design an algorithm and write code to find the
// first common ancestor of two nodes in a binary tree. Avoid storing additional
// nodes in a data structure. NOTE: This is not necessarily a binary search
// tree.
//
// Assumptions
//   * The root of the tree and the data of the target nodes is provided, but
//     the target nodes must be found within the tree.
//   * Nodes do not have knowledge of their parents.
//   * We may record that a node has been visited using a field on the node
//     itself.

// FirstCommonAncestor first checks that the first target node exists, and marks
// all nodes along the route to the target node as visited. It then searches for
// the second target node. Once found, it returns up the tree, checking if
// parent nodes are marked 'visited'. The first 'visited' node is the first
// common ancestor.
//
// Requires O(n) time, as the tree may not be a binary search tree. Requires
// O(log n) space on the stack for recursive calls.
func (n *BinaryTreeNode) FirstCommonAncestor(first, second int) *BinaryTreeNode {
	if !n.findAndMarkPath(first) {
		// First node not present in tree.
		return nil
	}

	return n.firstCommonAncestor(second)
}

func (n *BinaryTreeNode) findAndMarkPath(data int) bool {
	if n == nil {
		return false
	}

	// Node found; mark it as visited in case it is the common ancestor.
	if n.data == data {
		n.visited = true
		return true
	}

	// If either of the node's child branches contains the target node, mark it
	// as seen, tracing the path to the target node through the tree.
	if n.left.findAndMarkPath(data) || n.right.findAndMarkPath(data) {
		n.visited = true
		return true
	}

	return false
}

func (n *BinaryTreeNode) firstCommonAncestor(data int) *BinaryTreeNode {
	if n == nil {
		return nil
	}

	if n.data == data {
		return n
	}

	for _, child := range []*BinaryTreeNode{n.left, n.right} {
		ancestor := child.firstCommonAncestor(data)
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
