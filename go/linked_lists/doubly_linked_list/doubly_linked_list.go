package linkedlist

import "fmt"

type Node struct {
	Data       int
	Prev, Next *Node
}

func New(data int) *Node {
	return &Node{Data: data}
}

func (n *Node) AppendToTail(data int) {
	cur := n
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = &Node{Data: data, Prev: cur}
}

func (n *Node) AppendToHead(data int) *Node {
	cur := n
	for ; cur.Prev != nil; cur = cur.Prev {
	}
	cur.Prev = &Node{Data: data, Next: cur}
	return cur.Prev
}

func (n *Node) Delete(data int) *Node {
	cur := n
	if cur.Data == data {
		cur.Next.Prev = cur.Prev
		if cur.Prev != nil {
			cur.Prev.Next = cur.Next
		}
		return cur.Next
	}

	for ; cur.Next != nil; cur = cur.Next {
		if cur.Next.Data == data {
			cur.Next = cur.Next.Next
			if cur.Next != nil {
				cur.Next.Prev = cur
			}
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

func (n *Node) tail() *Node {
	cur := n
	for ; cur.Next != nil; cur = cur.Next {
	}
	return cur
}
