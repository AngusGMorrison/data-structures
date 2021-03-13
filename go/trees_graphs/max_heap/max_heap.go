// Package maxheap implements a max heap with integer data.
//
// A max heap is a complete binary tree in which the value of each parent node
// is greater than or equal to those of its children. It is represented as an
// array in which the root has index 0.
//
// Insertion and popping are O(logn) operations, as the heap property must
// be maintained by traversing up or down the heap, repectively, swapping nodes
// that violate the heap property.
package maxheap

import (
	"errors"
	"fmt"
)

type MaxHeap struct {
	heap []int
	size int
}

func New(capacity uint) *MaxHeap {
	return &MaxHeap{heap: make([]int, 0, capacity)}
}

// parent returns the index of the parent of the node at index idx.
func parent(idx int) int {
	return (idx - 1) / 2
}

// leftChild returns the index of the left child of the node at index idx.
func leftChild(idx int) int {
	return idx*2 + 1
}

// rightChild returns the index of the right child of the node at index idx.
func rightChild(idx int) int {
	return idx*2 + 2
}

// isLeaf returns true if the node is on the bottom row of the heap.
func (h *MaxHeap) isLeaf(idx int) bool {
	return idx >= ((h.size-1)/2) && idx < h.size
}

func (h *MaxHeap) swap(a, b int) {
	h.heap[a], h.heap[b] = h.heap[b], h.heap[a]
}

// Insert adds a new node to the heap, preserving the heap property. The new
// node is inserted at the bottom of the tree and is swapped upwards if greater
// than its parent.
func (h *MaxHeap) Insert(data int) error {
	if h == nil {
		return fmt.Errorf("can't insert %d into nil *MaxHeap", data)
	}

	h.heap = append(h.heap[:h.size], data)
	current := h.size
	prnt := parent(current)
	for h.heap[current] > h.heap[prnt] {
		h.swap(current, prnt)
		current = prnt
		prnt = parent(current)
	}
	h.size++

	return nil
}

// Pop removes and returns the root of the tree, preserving the heap property.
// The removed node is replaced by the last node in the tree, which is
// progessively swapped with its largest child node to restore the heap.
func (h *MaxHeap) Pop() (int, error) {
	if h == nil {
		return 0, errors.New("can't pop nil *MaxHeap")
	}

	popped := h.heap[0]
	h.size--
	h.heap[0] = h.heap[h.size]
	h.restore(0)

	return popped, nil
}

// restore recurses down the tree starting at node with index idx. It swaps any
// parent node smaller than at least one of its child nodes with the greatest
// child node until the heap property is restored.
func (h *MaxHeap) restore(idx int) {
	current := h.heap[idx]
	leftIdx, rightIdx := leftChild(idx), rightChild(idx)

	if h.isLeaf(idx) || current > h.heap[leftIdx] && current > h.heap[rightIdx] {
		return
	}

	if h.heap[leftIdx] > h.heap[rightIdx] {
		h.swap(idx, leftIdx)
		h.restore(leftIdx)
	} else {
		h.swap(idx, rightIdx)
		h.restore(rightIdx)
	}
}

func (h *MaxHeap) Print() {
	h.print(0, 0)
}

func (h *MaxHeap) print(idx, depth int) {
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
