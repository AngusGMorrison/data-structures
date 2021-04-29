package ctci

// 4.3 List of Depths: Given a binary tree, design an algorithm which creates a
// linked list of all the nodes at each depth (e.g. if you have a tree with
// depth D, you'll have D linked lists).

// GenerateDepthListsRecursive returns a map of depths to linked lists
// containing all nodes in the input tree at that depth. The lists are built
// recursively by incrementing the depth as it is passed down the call stack.
//
// Takes O(n) time and O(n) space – we must visit each node and we must store
// each tree node in a list node.
//
// Assumptions:
//   * A map of linked lists is an acceptable return value. Confirm this with
//     the interviewer.
//   * The tree nodes must be wrapped in linked list nodes. I.e. the tree nodes
//     do not have next fields.
//   * Root node is at depth 1.
//   * The order in which nodes are appended/prepended to the linked lists is unimportant.
func GenerateDepthListsRecursive(root *BinaryTreeNode) map[int]*BinaryTreeListNode {
	if root == nil {
		return nil
	}

	depthLists := make(map[int]*BinaryTreeListNode)
	listBuilder := func(depth int, n *BinaryTreeNode) {
		depthLists[depth] = NewBinaryTreeListNode(n, depthLists[depth])
	}
	root.MapWithDepth(1, listBuilder)

	return depthLists
}

// GenerateDepthListsBFS returns a map of depths to linked lists containing all
// nodes in the input tree at that depth. The lists are built by breadth-first
// search (BFS), incrementing the depth each time the search moves one level
// deeper.
// Takes O(n) time and O(n) space. Although BFS doesn't require log n
// recursive calls on the stack, it must still create a linked list node for
// each node in the input tree.
func GenerateDepthListsBFS(root *BinaryTreeNode) map[int]*BinaryTreeListNode {
	if root == nil {
		return nil
	}

	depthLists := make(map[int]*BinaryTreeListNode)
	worklist := []*BinaryTreeNode{root}
	var depth int

	for len(worklist) > 0 {
		nodes := worklist
		worklist = nil
		depth += 1

		for _, n := range nodes {
			depthLists[depth] = NewBinaryTreeListNode(n, depthLists[depth])

			if n.left != nil {
				worklist = append(worklist, n.left)
			}
			if n.right != nil {
				worklist = append(worklist, n.right)
			}
		}
	}

	return depthLists
}
