package ctci

// 2.1 Remove Dups: Write code to remove duplicates from an unsorted linked
// list.
//
// FOLLOW UP: How would you solve this problem if a temporary buffer is not
// allowed?
//
// Assumptions:
//	* Confirm whether list is singly or doubly linked.
//  * For similar questions, confirm whether list is sorted.

// DedupSingle checks for duplicates in a singly linked list by storing the data
// from each seen node in a set and checking subsequent nodes against those that
// have already been seen.
//
// Time: O(n)
// Space: O(n)
func DedupSingle(head *SNode) {
	seen := make(map[int]bool)
	var prev *SNode
	for cur := head; cur != nil; cur = cur.Next {
		if seen[cur.Data] {
			prev.Next = cur.Next
		} else {
			seen[cur.Data] = true
			prev = cur // as head is never a dup, prev is always assigned before it is accessed
		}
	}
}

// DedupDouble checks for duplicates in a doubly linked list by storing the data
// from each seen node in a set and checking subsequent nodes against those that
// have already been seen.
//
// Time: O(n)
// Space: O(n)
func DedupDouble(head *DNode) {
	seen := make(map[int]bool)
	var prev *DNode
	for cur := head; cur != nil; cur = cur.Next {
		if seen[cur.Data] {
			prev.Next = cur.Next
			if cur.Next != nil {
				cur.Next.Prev = prev
			}
		} else {
			seen[cur.Data] = true
			prev = cur
		}
	}
}

// DedupSingleNoBuffer checks for duplicates by visiting each node in the
// singly linked list in turn and checking the rest of the list for matches.
//
// Time: O(n^2)
// Space: O(1)
func DedupSingleNoBuffer(head *SNode) {
	for uniq := head; uniq != nil; uniq = uniq.Next {
		for cur := uniq; cur.Next != nil; cur = cur.Next {
			if cur.Next.Data == uniq.Data {
				cur.Next = cur.Next.Next
			}
		}
	}
}

// DedupDoubleNoBuffer checks for duplicates by visiting each node in the
// doubly linked list in turn and checking the rest of the list for matches.
//
// Time: O(n^2)
// Space: O(1)
func DedupDoubleNoBuffer(head *DNode) {
	for uniq := head; uniq != nil; uniq = uniq.Next {
		for cur := uniq; cur.Next != nil; cur = cur.Next {
			if cur.Next.Data == uniq.Data {
				cur.Next = cur.Next.Next
				if cur.Next != nil {
					cur.Next.Prev = cur
				}
			}
		}
	}
}
