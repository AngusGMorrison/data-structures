package ctci

// 4.9 BST Sequences: A binary search tree was created by traversing through an
// array from left to right and inserting each element. Given a binary search
// tree with distinct elements, print all possible arrays that could have led to
// this tree.

// The next node in a binary tree sequence may be any one of the child nodes of
// the nodes currently in the sequence which have no been fully explored.
//
// Like a breadth-first search, we keep a worklist of nodes which have not yet
// been added to the permutations. Each time a node is processed, its child
// nodes are added to the worklist.
func (n *BinaryTreeNode) Sequences() [][]int {
	if n == nil {
		return nil
	}

	permutations := make([][]int, 0)

	var permute func([]*BinaryTreeNode, []int)
	permute = func(worklist []*BinaryTreeNode, perm []int) {
		if len(worklist) == 0 {
			permutations = append(permutations, perm)
			return
		}

		for i, node := range worklist {
			nextList := nextWorklist(worklist, i)
			nextPerm := nextPermutation(perm, node)

			permute(nextList, nextPerm)
		}
	}

	permute([]*BinaryTreeNode{n}, []int{})

	return permutations
}

// nextWorklist generates the next set of valid sequence nodes by copying the
// current worklist, extracting the current node at idx and appending its
// children to the new list.
func nextWorklist(worklist []*BinaryTreeNode, idx int) []*BinaryTreeNode {
	currentNode := worklist[idx]
	nextList := make([]*BinaryTreeNode, len(worklist))
	copy(nextList, worklist)
	nextList = append(nextList[:idx], nextList[idx+1:]...)
	nextList = append(nextList, currentNode.Children()...)

	return nextList
}

// nextPermutation extends an unfinished permutation by appending the current
// node's data to a copy of the slice.
func nextPermutation(perm []int, n *BinaryTreeNode) []int {
	next := make([]int, len(perm), len(perm)+1)
	copy(next, perm)
	next = append(next, n.data)

	return next
}
