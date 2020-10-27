package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func New(data int) *Node {
	return &Node{Data: data}
}

func (n *Node) AppendToHead(data int) *Node {
	new := New(data)
	new.Next = n
	return new
}

func (n *Node) AppendToTail(data int) {
	cur := n
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = New(data)
}

func (n *Node) Delete(data int) *Node {
	cur := n
	if cur.Data == data {
		return cur.Next // head shift
	}

	for ; cur.Next != nil; cur = cur.Next {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
			return n
		}
	}

	return n
}

func (n *Node) GetIndex(idx int) (int, error) {
	cur := n
	for i := idx; i > 0 && cur != nil; i-- {
		cur = cur.Next
	}

	if cur == nil {
		return -1, fmt.Errorf("index %d out of bounds", idx)
	}
	return cur.Data, nil
}
