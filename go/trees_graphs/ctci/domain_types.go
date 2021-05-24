package ctci

import (
	"strconv"
	"strings"
)

// DirectedGraphNode represents a node in a directed graph.
type DirectedGraphNode struct {
	name     string
	children []*DirectedGraphNode
}

func NewDirectedGraphNode(name string) *DirectedGraphNode {
	return &DirectedGraphNode{
		name:     name,
		children: make([]*DirectedGraphNode, 0),
	}
}

// BidirecGraphNode is a directed node which maintains knowledge of its parents.
type BidirecGraphNode struct {
	name     string
	children []*BidirecGraphNode
	parents  []*BidirecGraphNode
}

func NewBirdirecGraphNode(name string) *BidirecGraphNode {
	return &BidirecGraphNode{
		name:     name,
		children: make([]*BidirecGraphNode, 0),
		parents:  make([]*BidirecGraphNode, 0),
	}
}

type BinaryTreeNode struct {
	data                int
	left, right, parent *BinaryTreeNode
	visited             bool
	size                int
}

func NewBinaryTreeNode(data int, left, right, parent *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{
		data:    data,
		left:    left,
		right:   right,
		parent:  parent,
		visited: false,
	}
}

// Insert inserts a node into the binary tree. Note that the binary tree is not balanced.
func (n *BinaryTreeNode) Insert(data int) *BinaryTreeNode {
	return n.insert(data, nil)
}

func (n *BinaryTreeNode) insert(data int, parent *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return NewBinaryTreeNode(data, nil, nil, parent)
	}

	if data <= n.data {
		n.left = n.left.insert(data, n)
	} else {
		n.right = n.right.insert(data, n)
	}

	return n
}

// MassInsert inserts one node into the binary tree for each argument passed.
func (n *BinaryTreeNode) MassInsert(data ...int) *BinaryTreeNode {
	for _, d := range data {
		n = n.Insert(d)
	}

	return n
}

// Each iterates over each node in the binary tree in order, applying the
// provided function.
func (n *BinaryTreeNode) Each(fn func(n *BinaryTreeNode)) {
	if n == nil {
		return
	}

	n.left.Each(fn)
	fn(n)
	n.right.Each(fn)
}

// EachWithDepth iterates over each node in a binary tree in order, incrementing
// the starting depth to reflect the current level of the tree, and applies the
// provided function to the node.
func (n *BinaryTreeNode) EachWithDepth(depth int, fn func(depth int, n *BinaryTreeNode)) {
	if n == nil {
		return
	}

	n.left.EachWithDepth(depth+1, fn)
	fn(depth, n)
	n.right.EachWithDepth(depth+1, fn)
}

func (n *BinaryTreeNode) Sibling() *BinaryTreeNode {
	if n == nil || n.parent == nil {
		return nil
	}

	if n.parent.left == n {
		return n.parent.right
	}

	return n.parent.left
}

// Children returns a slice containing the children of a BinaryTreeNode, if
// present.
func (n *BinaryTreeNode) Children() []*BinaryTreeNode {
	if n == nil {
		return nil
	}

	children := make([]*BinaryTreeNode, 0, 2)
	if n.left != nil {
		children = append(children, n.left)
	}
	if n.right != nil {
		children = append(children, n.right)
	}

	return children
}

func (n *BinaryTreeNode) Size() int {
	if n == nil {
		return 0
	}

	if n.size == 0 {
		leftSize := n.left.Size()
		rightSize := n.right.Size()

		n.size = leftSize + rightSize + 1
	}

	return n.size
}

func (n *BinaryTreeNode) String() string {
	return strconv.Itoa(n.data)
}

// BinaryTreeListNode is a linked list node which holds a binary tree node as
// data.
type BinaryTreeListNode struct {
	treeNode *BinaryTreeNode
	next     *BinaryTreeListNode
}

func NewBinaryTreeListNode(treeNode *BinaryTreeNode, next *BinaryTreeListNode) *BinaryTreeListNode {
	return &BinaryTreeListNode{
		treeNode: treeNode,
		next:     next,
	}
}

func (ln *BinaryTreeListNode) String() string {
	var sb strings.Builder
	for cur := ln; ln != nil; ln = ln.next {
		sb.WriteString(strconv.Itoa(cur.treeNode.data))
		if ln.next != nil {
			sb.WriteString(" -> ")
		}
	}

	return sb.String()
}

type AdjacencyList map[string][]string
