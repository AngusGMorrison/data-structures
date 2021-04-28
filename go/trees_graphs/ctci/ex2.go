package ctci

import (
	"strconv"
	"strings"
)

// 4.2 Given a sorted (increasing order) array with unique integer elements,
// write an algorithm to create a binary search tree with minimal height.

// TreeFromSlice builds a tree recursively from a sorted slice. At each level,
// the middle node of the slice is chosen as the root. Elements to the left of
// the root become the left branch, and elements to the right of the root become
// the right branch.
func TreeFromSlice(arr []int) *BinaryTreeNode {
	if len(arr) == 0 {
		return nil
	}

	rootIdx := len(arr) / 2
	root := &BinaryTreeNode{data: arr[rootIdx]}
	root.left = TreeFromSlice(arr[:rootIdx])
	root.right = TreeFromSlice(arr[rootIdx+1:])
	return root
}

func (tn *BinaryTreeNode) String() string {
	var sb strings.Builder
	tn.string(0, &sb)
	return sb.String()
}

func (tn *BinaryTreeNode) string(depth int, sb *strings.Builder) {
	if tn == nil {
		return
	}

	tn.right.string(depth+1, sb)

	for i := 0; i < depth; i++ {
		sb.WriteByte('\t')
	}
	sb.WriteString(strconv.Itoa(tn.data))
	sb.WriteByte('\n')

	tn.left.string(depth+1, sb)
}
