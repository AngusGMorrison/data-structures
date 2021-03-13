// Package minheap implements a min heap that stores integers.
//
// A min heap is a complete binary tree in which the value of each parent node
// is greater than that of its children. It is represented as an array in which
// the root has index 0.
//
// Insertion and deletion are O(logn) operations, as the heap property must be
// maintained by traversing up or down the heap, respectively, and swapping
// nodes which violate the heap property.
package minheap

import (
	"errors"
	"fmt"
)

type MinHeap struct {
	heap []int
	size int
}

func New(capacity uint) *MinHeap {
	return &MinHeap{heap: make([]int, 0, capacity)}
}

// parent returns the index of the parent of the node with index idx.
func parent(idx int) int {
	return (idx - 1) / 2
}

// leftChild returns the index of the left child of the node with index idx.
func leftChild(idx int) int {
	return idx*2 + 1
}

// rightChild returns the index of the right child of the node with index idx.
func rightChild(idx int) int {
	return idx*2 + 2
}

// isLeaf returns true if the node is on the bottom row of the tree.
func (h *MinHeap) isLeaf(idx int) bool {
	return idx >= ((h.size-1)/2) && idx < h.size
}

func (h *MinHeap) swap(a, b int) {
	h.heap[a], h.heap[b] = h.heap[b], h.heap[a]
}

func (h *MinHeap) Insert(data int) error {
	if h == nil {
		return errors.New("can't insert %d into nil *MinHeap")
	}

	h.heap = append(h.heap[:h.size], data)
	current := h.size // the index of the current node as we've yet to increment size
	prnt := parent(current)
	for h.heap[current] < h.heap[prnt] {
		h.swap(current, prnt)
		current = prnt
		prnt = parent(current)
	}

	h.size += 1
	return nil
}

func (h *MinHeap) Pop() (int, error) {
	if h == nil {
		return 0, errors.New("can't delete from nil *MinHeap")
	}

	popped := h.heap[0]
	h.size -= 1
	// Move the last element of the heap to the root and sift downwards to
	// restore the heap property.
	h.heap[0] = h.heap[h.size]
	h.restore(0)
	return popped, nil
}

// restore recurses down the tree starting at node with index idx. It swaps any
// parent node greater than at least one of its child nodes with the smallest
// child node until the heap property is restored.
func (h *MinHeap) restore(idx int) {
	current := h.heap[idx]
	leftIdx, rightIdx := leftChild(idx), rightChild(idx)

	// When a leaf is reached or the current node is smaller than both of its
	// children, the heap is restored.
	if h.isLeaf(idx) || current < h.heap[leftIdx] && current < h.heap[rightIdx] {
		return
	}

	if h.heap[leftIdx] < h.heap[rightIdx] {
		h.swap(idx, leftIdx)
		h.restore(leftIdx)
	} else {
		h.swap(idx, rightIdx)
		h.restore(rightIdx)
	}
}

func (h *MinHeap) Print() {
	h.print(0, 0)
}

func (h *MinHeap) print(idx, depth int) {
	if idx >= h.size {
		return
	}

	h.print(rightChild(idx), depth+1)
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}
	fmt.Println(h.heap[idx])
	h.print(leftChild(idx), depth+1)
}
