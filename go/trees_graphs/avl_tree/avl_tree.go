// Package avl implements an AVL tree, a balanced binary search tree that
// enables search, insertion and deletion in O(logn) time.
//
// Better balanced than red-black trees, AVL trees require more rotations to
// remain in balance, and so are more suited for use cases where lookup occurs
// more often than modification.
package avl

import "fmt"

// Node is a simple tree node which stores and sorts on an integer value.
type Node struct {
	data, height int
	left, right  *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Height returns the node's height.
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// Data returns the node's data.
func (n *Node) Data() int {
	return n.data
}

func (n *Node) updateHeight() {
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

func (n *Node) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

func (n *Node) inOrderSuccessor() *Node {
	if n == nil {
		return nil
	}

	return n.right.minNode()
}

func (n *Node) minNode() *Node {
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}
	return n
}

func newNode(data int) *Node {
	return &Node{
		data:   data,
		height: 1,
	}
}

// Insert accepts an integer and returns the root of the balanced tree containing a node with the
// given value. Duplicates are not permitted. The returned boolean is false if the value already
// exists in the tree.
func (n *Node) Insert(data int) (*Node, bool) {
	if n == nil {
		return newNode(data), true
	}

	if data < n.data {
		left, ok := n.left.Insert(data)
		if !ok {
			return n, false
		}
		n.left = left
	} else if data > n.data {
		right, ok := n.right.Insert(data)
		if !ok {
			return n, false
		}
		n.right = right
	} else {
		// Duplicate data; do nothing.
		return n, false
	}

	return n.rotate(), true
}

// Delete accepts an integer and returns the root of the balanced tree from which the node with that
// value has been deleted. The returned boolean is false if the value does not exist in the tree.
func (n *Node) Delete(data int) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	if data < n.data {
		left, ok := n.left.Delete(data)
		if !ok {
			return n, false
		}
		n.left = left
	} else if data > n.data {
		right, ok := n.right.Delete(data)
		if !ok {
			return n, false
		}
		n.right = right
	} else {
		n = deleteNode(n)
	}

	return n.rotate(), true
}

func deleteNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.left != nil && n.right != nil {
		successor := n.inOrderSuccessor()
		n.data = successor.data
		n.right, _ = n.right.Delete(successor.data)
		return n
	}

	var child *Node
	if n.left != nil {
		child = n.left
	} else {
		child = n.right
	}

	return child
}

func (n *Node) rotate() *Node {
	if n == nil {
		return nil
	}

	n.updateHeight()
	bFactor := n.balanceFactor()

	if bFactor > 1 {
		// Left-heavy.
		lChildBFactor := n.left.balanceFactor()
		if lChildBFactor >= 0 {
			// Long left chain.
			return n.rotateRight()
		}

		// Long right chain.
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	}

	if bFactor < -1 {
		// Right-heavy.
		rChildBFactor := n.right.balanceFactor()
		if rChildBFactor <= 0 {
			// Long right chain.
			return n.rotateLeft()
		}

		// Long left chain.
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func (n *Node) rotateRight() *Node {
	newRoot := n.left
	newNLeft := newRoot.right
	newRoot.right = n
	n.left = newNLeft

	// Note the order of height updates.
	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

func (n *Node) rotateLeft() *Node {
	newRoot := n.right
	newNRight := newRoot.left
	newRoot.left = n
	n.right = newNRight

	// Note the order of height updates.
	n.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// Print executes a post-order traversal of the tree with root n, so that the
// greatest nodes will be printed first and the least nodes last. Each node is
// indented by a number of tabs equal to its depth in the tree.
func (n *Node) Print(depth int) {
	if n == nil {
		return
	}

	n.right.Print(depth + 1)
	for tabs := depth; tabs > 0; tabs-- {
		fmt.Print("\t")
	}
	fmt.Println("data: ", n.data, "height: ", n.Height())
	n.left.Print(depth + 1)
}
