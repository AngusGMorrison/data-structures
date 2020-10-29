package ctci

import "testing"

func TestDeleteMiddle(t *testing.T) {
	t.Run("deletes a node in the middle of a linked list", func(t *testing.T) {
		length := 3
		head := nNodeSinglyLinkedList(uint(length))
		valToDelete := 1
		err := DeleteMiddle(head.Next)
		if err != nil {
			t.Errorf("got error %v", err)
		}

		cur := head
		for i := 0; i < length-1; i++ {
			if cur == nil || cur.Data == 1 {
				t.Errorf("want node to exist and not have data %d, got %+v", valToDelete, cur)
			}
			cur = cur.Next
		}
		if cur != nil {
			t.Errorf("want list to have length %d, got extra node %+v", length-1, cur)
		}
	})

	t.Run("returns an error if passed a tail node", func(t *testing.T) {
		head := nNodeSinglyLinkedList(2)
		err := DeleteMiddle(head.Next)
		if err == nil {
			t.Errorf("want error, got nil")
		}
	})
}
