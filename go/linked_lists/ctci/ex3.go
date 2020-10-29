package ctci

import "errors"

// 1.3 Delete the Middle Node: Implement an algorithm to delete a node in the
// middle (i.e. any node but the first and last node) of a singly linked list,
// given access only to that node.

// DeleteMiddle "deletes" a middle node by copying the data of its next node
// into itself and deleting the next node.
//
// Time complexity: O(1)
// Space complexity: O(1)
func DeleteMiddle(s *SNode) error {
	if s == nil || s.Next == nil {
		return errors.New("s cannot be tail node or nil")
	}

	next := s.Next
	s.Data = next.Data
	s.Next = next.Next
	return nil
}
