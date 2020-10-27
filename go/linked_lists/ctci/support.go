package ctci

import "fmt"

type SNode struct {
	Data int
	Next *SNode
}

type DNode struct {
	Data       int
	Prev, Next *DNode
}

func newSinglyLinkedList(data int) *SNode {
	return &SNode{Data: data}
}

func newDoublyLinkedList(data int) *DNode {
	return &DNode{Data: data}
}

func (s *SNode) appendToTail(data int) {
	cur := s
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = newSinglyLinkedList(data)
}

func (d *DNode) appendToTail(data int) {
	cur := d
	for ; cur.Next != nil; cur = cur.Next {
	}
	cur.Next = newDoublyLinkedList(data)
	cur.Next.Prev = cur
}

func (s *SNode) appendToHead(data int) *SNode {
	new := newSinglyLinkedList(data)
	new.Next = s
	return new
}

func (d *DNode) appendToHead(data int) *DNode {
	d.Prev = newDoublyLinkedList(data)
	d.Prev.Next = d
	return d.Prev
}

func nNodeSinglyLinkedList(n uint) *SNode {
	list := newSinglyLinkedList(int(n - 1))
	for i := int(n - 2); i >= 0; i-- {
		list = list.appendToHead(i)
	}
	return list
}

func nNodeDoublyLinkedList(n uint) *DNode {
	list := newDoublyLinkedList(int(n - 1))
	for i := int(n - 2); i >= 0; i-- {
		list = list.appendToHead(i)
	}
	return list
}

func (s *SNode) len() int {
	var count int
	for cur := s; cur != nil; cur = cur.Next {
		count++
	}
	return count
}

func (d *DNode) len() int {
	var count int
	for cur := d; cur != nil; cur = cur.Next {
		count++
	}
	return count
}

func (s *SNode) print() {
	data := make([]int, 0, s.len())
	for cur := s; cur != nil; cur = cur.Next {
		data = append(data, cur.Data)
	}
	fmt.Println(data)
}

func (d *DNode) print() {
	data := make([]int, 0, d.len())
	for cur := d; cur != nil; cur = cur.Next {
		data = append(data, cur.Data)
	}
	fmt.Println(data)
}
