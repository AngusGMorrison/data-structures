package ctci

import (
	"math/rand"
	"time"
)

// 4.11 Random Node: You are implementing a binary tree class from scratch
// which, in addition to insert, find, and delete, has a method getRandomNode()
// which returns a random node from the tree. All nodes should be equally likely
// to be chosen. Design and implement an algorithm for getRandomNode, and
// explain how you would implement the rest of the methods.

// Assumptions:
//   * We are implementing a binary search tree.

type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Size() int {
	return bst.root.size
}

type BSTNode struct {
	data        int
	left, right *BSTNode
	size        int
}

func (n *BSTNode) Data() int {
	if n == nil {
		return 0
	}
	return n.data
}

func (n *BSTNode) Left() *BSTNode {
	if n == nil {
		return nil
	}

	return n.left
}

func (n *BSTNode) Right() *BSTNode {
	if n == nil {
		return nil
	}

	return n.right
}

func (bst *BST) Insert(data int) *BSTNode {
	if bst == nil {
		return nil
	}

	insertion := &BSTNode{data: data, size: 1}

	if bst.root == nil {
		bst.root = insertion
	} else {
		insert(bst.root, insertion)
	}

	return insertion
}

func insert(root, insertion *BSTNode) *BSTNode {
	if root == nil {
		return insertion
	}

	if insertion.data <= root.data {
		root.left = insert(root.left, insertion)
	} else {
		root.right = insert(root.right, insertion)
	}

	root.size++
	return root
}

func (bst *BST) Find(data int) *BSTNode {
	if bst == nil || bst.root == nil {
		return nil
	}

	return find(bst.root, data)
}

func find(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}

	if data < root.data {
		return find(root.left, data)
	} else if data > root.data {
		return find(root.right, data)
	}
	return root
}

// Delete removes a node from the tree by replacing it with: nil, if it is
// a leaf node; its child, if it has only one child; its in-order successor, if
// it has two children.
func (bst *BST) Delete(data int) bool {
	if bst == nil || bst.root == nil {
		return false
	}

	root, deleted := deleteNode(bst.root, data)
	if deleted {
		bst.root = root
	}

	return deleted
}

func deleteNode(root *BSTNode, data int) (n *BSTNode, deleted bool) {
	if root == nil {
		return nil, false
	}

	if data < root.data {
		root.left, deleted = deleteNode(root.left, data)
	} else if data > root.data {
		root.right, deleted = deleteNode(root.right, data)
	} else {
		leftChild := root.left
		rightChild := root.right
		if leftChild == nil && rightChild == nil {
			return nil, true
		} else if leftChild == nil {
			return rightChild, true
		} else if rightChild == nil {
			return leftChild, true
		}

		return deleteTwoChildNode(root), true
	}

	if deleted {
		root.size--
	}

	return root, deleted
}

func deleteTwoChildNode(root *BSTNode) *BSTNode {
	leftChild := root.left
	rightChild := root.right

	if leftChild == nil && rightChild == nil {
		return nil
	}
	if leftChild == nil {
		return rightChild
	}
	if rightChild == nil {
		return leftChild
	}

	inOrderSuccessor := inOrderSuccessor(root)
	inOrderSuccessor.left = leftChild
	inOrderSuccessor.right = rightChild
	return inOrderSuccessor
}

func inOrderSuccessor(root *BSTNode) *BSTNode {
	var successor *BSTNode
	for successor := root.right; successor.left != nil; successor = successor.left {
	}
	return successor
}

// GetRandomNodeSequential returns a random node by generating a random node
// number between 1 and the tree's size, and iterating over the tree via
// pre-order traversal. With each additional node visited, a counter is
// increased by 1. When the counter and the node number are equal, the current
// node is returned.
//
// Requires O(n) time and O(d) space, where d is the depth of the tree.
func (bst *BST) GetRandomNodeSequential() *BSTNode {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	target := rng.Intn(bst.Size())
	var counter int
	var randomNode *BSTNode

	var traverse func(n *BSTNode)
	traverse = func(n *BSTNode) {
		if n == nil {
			return
		}

		if counter == target {
			randomNode = n
			return
		}
		counter++

		traverse(n.left)
		traverse(n.right)
	}

	traverse(bst.root)

	return randomNode
}

// GetRandomNodeLogarithmic returns a random node by generating a random index
// between 0 and the size of the tree, and then recusively moving to the half of
// the tree where the node with that index must reside. We are locating the ith
// node in an in-order traversal. This requires each node to know the size of
// its subtree.
//
// Takes O(d) time and O(d) space, where d is the max depth of the tree.
func (bst *BST) GetRandomNodeLogarithmic() *BSTNode {
	if bst.root == nil {
		return nil
	}

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	idx := rng.Intn(bst.Size())

	return getNodeAtIdx(bst.root, idx)
}

func getNodeAtIdx(n *BSTNode, idx int) *BSTNode {
	var leftSize int
	if n.left != nil {
		leftSize = n.left.size
	}

	if idx < leftSize {
		return getNodeAtIdx(n.left, idx)
	} else if idx > leftSize {
		// Skip over the left half of the tree, including the current node.
		return getNodeAtIdx(n.right, idx-(leftSize+1))
	} else {
		return n
	}
}
