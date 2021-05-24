package ctci

import (
	"container/list"
)

// 4.9 BST Sequences: A binary search tree was created by traversing through an
// array from left to right and inserting each element. Given a binary search
// tree with distinct elements, print all possible arrays that could have led to
// this tree.

func (n *BinaryTreeNode) Sequences() []*list.List {
	result := []*list.List{}

	if n == nil {
		result = append(result, list.New())
		return result
	}

	// Capture the current node, which must appear before its left and right
	// subtrees in every permuation.
	permutation := list.New()
	permutation.PushBack(n.data)

	// Recurse on left and right subtrees.
	leftSeqs := n.left.Sequences()
	rightSeqs := n.right.Sequences()

	// Weave all lists from the left and right sides together, prefixed with the
	// current node.
	for _, leftList := range leftSeqs {
		for _, rightList := range rightSeqs {
			permutations := []*list.List{}
			weave(leftList, rightList, permutation, &permutations)
			result = append(result, permutations...)
		}
	}

	return result
}

// weave produces all possible permutations of two linked lists by recursively
// removing the head of the left list and appending it to the current
// permutation. When the left list is empty, the remainder of the second list is
// appended. This process is then repeated with the right list.
func weave(left, right, permutation *list.List, results *[]*list.List) {
	// If one list is empty, clone the permutation and add the remainder of the
	// other list to complete it, then store the result.
	if left.Len() == 0 || right.Len() == 0 {
		result := clone(permutation)
		pushBackAll(result, left)
		pushBackAll(result, right)
		*results = append(*results, result)
		return
	}

	// Recurse with the head of left added to the permuation. Removing the head
	// will damage left, so we need to put it back where we found it afterwards.
	leftHead := left.Remove(left.Front()).(int)
	permutation.PushBack(leftHead)
	weave(left, right, permutation, results)
	permutation.Remove(permutation.Back())
	left.PushFront(leftHead)

	// Do the same thing with the second, damaging and then restoring the list.
	rightHead := right.Remove(right.Front()).(int)
	permutation.PushBack(rightHead)
	weave(left, right, permutation, results)
	permutation.Remove(permutation.Back())
	right.PushFront(rightHead)
}

func pushBackAll(dest, src *list.List) {
	for cur := src.Front(); cur != nil; cur = cur.Next() {
		dest.PushBack(cur.Value)
	}
}

func clone(l *list.List) *list.List {
	dup := list.New()
	for cur := l.Front(); cur != nil; cur = cur.Next() {
		dup.PushBack(cur.Value)
	}

	return dup
}
