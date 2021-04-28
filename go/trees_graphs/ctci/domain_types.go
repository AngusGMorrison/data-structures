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
	data        int
	left, right *BinaryTreeNode
}

func NewBinaryTreeNode(data int, left, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{
		data:  data,
		left:  left,
		right: right,
	}
}

// Insert inserts a node into the binary tree. Note that the binary tree is not balanced.
func (n *BinaryTreeNode) Insert(data int) *BinaryTreeNode {
	if n == nil {
		return NewBinaryTreeNode(data, nil, nil)
	}

	if data < n.data {
		n.left = n.left.Insert(data)
	} else if data > n.data {
		n.right = n.right.Insert(data)
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

// MapWithDepth iterates over each node in a binary tree in order, incrementing
// the starting depth to reflect the current level of the tree, and applies the
// provided function to the node.
func (n *BinaryTreeNode) MapWithDepth(depth int, fn func(depth int, n *BinaryTreeNode)) {
	if n == nil {
		return
	}

	n.left.MapWithDepth(depth+1, fn)
	fn(depth, n)
	n.right.MapWithDepth(depth+1, fn)
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
