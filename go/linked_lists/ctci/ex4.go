package ctci

// 1.4 Partition: Write code to partition a linked list around a value x, such
// that all the nodes less than x come before all nodes greater than or equal
// to x. If x is contained within the list, the values of x only need to be
// after the elements less than x. The partition element x can appear anywhere
// in the right partition; it does not need to appear between th left and right
// partitions.
//
// EXAMPLE
// Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition = 5]
// Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
//
// Assumptions:
//	* Singly linked list
//  * Confirm whether partition should be stable or unstable
//	* List does not contain cycles

// PartitionStable constructs two linked lists, one containing elements less
// than val, and the other containing all other elements, then joins the lists
// together, preserving the relative ordering of of elements in each partition.
//
// Time complexity: O(n)
// Space complexity: O(1)
func PartitionStable(s *SNode, val int) *SNode {
	var firstHalfHead, firstHalfTail, secondHalfHead, secondHalfTail *SNode
	for cur := s; cur != nil; {
		// Setting each cur.Next to nil prevents needings to check each half
		// for nil and setting the appropriate tail to nil when the loop exits.
		next := cur.Next
		cur.Next = nil
		if cur.Data < val {
			appendToTail(firstHalfHead, firstHalfTail, cur)
		} else {
			appendToTail(secondHalfHead, secondHalfTail, cur)
		}
		cur = next
	}

	// Ensure first half is not nil before attempting to join lists together.
	if firstHalfHead == nil {
		return secondHalfHead
	}

	firstHalfTail.Next = secondHalfHead
	return firstHalfHead
}

func appendToTail(head, tail, n *SNode) {
	if head == nil {
		head = n
		tail = n
	} else {
		tail.Next = n
		tail = n
	}
}

// PartitionUnstable appends to the head of a new list if the node's value is
// less than the partition value, and to the tail if it is equal to or greater.
//
// Time complexity: O(n)
// Space complexity: O(1)
func PartitionUnstable(n *SNode, val int) *SNode {
	head := n
	tail := head
	for cur := head.Next; cur != nil; {
		next := cur.Next
		if cur.Data < val {
			cur.Next = head
			head = cur
		} else {
			tail.Next = cur
			tail = cur
		}

		cur = next
	}
	tail.Next = nil

	return head
}
