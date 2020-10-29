package ctci

import "errors"

// 1.2 Return Kth To Last: Implement an algorithm to find the kth to last
// element of a singly linked list.
//
// Assumptions:
//	* Function should return the node's data, not the node itself
//	* k == 1 returns the last node in the list
//	* The length of the list is not given.

// KthToLastIterative moves a fast runner k steps into the linked list before
// introducing a slow runner at the head of the list. It then moves the runners
// together until the fast runner reaches the end of the list, at which point
// the slow runner is k steps from the end of the list.
//
// Time complexity: O(n)
// Space complexity: O(1)
func (s *SNode) KthToLastIterative(k uint) (int, error) {
	if k == 0 {
		return 0, ErrKOutOfBounds
	}

	lead := s
	for i := uint(0); i < k; i++ {
		if lead == nil {
			return 0, ErrKOutOfBounds
		}
		lead = lead.Next
	}

	follow := s
	for lead != nil {
		lead = lead.Next
		follow = follow.Next
	}

	return follow.Data, nil
}

// KthToLastRecursive recurses to the end of the list and increments a counter
// as it returns. When the counter == k, the correct node has been identified
// and it is passed up the call stack.
//
// Time complexity: O(n)
// Space complexity: O(n)
func (s *SNode) KthToLastRecursive(k uint) (int, error) {
	var index uint
	var kthToLast func(s *SNode) *SNode
	kthToLast = func(s *SNode) *SNode {
		if s == nil {
			return nil
		}

		next := kthToLast(s.Next)
		index++
		if index == k {
			return s
		}
		return next
	}

	kthLastNode := kthToLast(s)
	if kthLastNode == nil {
		return 0, ErrKOutOfBounds
	}
	return kthLastNode.Data, nil
}

var ErrKOutOfBounds = errors.New("k is out of bounds")
