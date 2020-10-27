package ctci

import "testing"

func TestDedupSingle(t *testing.T) {
	impls := []struct {
		name string
		fn   func(*SNode)
	}{
		{"DedupSingle", DedupSingle},
		{"DedupSingleNoBuffer", DedupSingleNoBuffer},
	}

	for _, impl := range impls {
		t.Run(impl.name, func(t *testing.T) {
			list := buildDupListSingle()
			impl.fn(list)
			if gotLen := list.len(); gotLen != testListLen {
				t.Errorf("want len %d, got %d", testListLen, gotLen)
			}
		})
	}
}

func TestDedupDouble(t *testing.T) {
	impls := []struct {
		name string
		fn   func(*DNode)
	}{
		{"DedupDouble", DedupDouble},
		{"DedupDoubleNoBuffer", DedupDoubleNoBuffer},
	}

	for _, impl := range impls {
		t.Run(impl.name, func(t *testing.T) {
			list := buildDupListDouble()
			DedupDouble(list)
			if gotLen := list.len(); gotLen != testListLen {
				t.Errorf("want len %d, got %d", testListLen, gotLen)
			}
		})
	}
}

const testListLen = 3

func buildDupListSingle() *SNode {
	list := newSinglyLinkedList(0)
	for i := 0; i < testListLen; i++ {
		list.appendToTail(i)
	}

	return list
}

func buildDupListDouble() *DNode {
	list := newDoublyLinkedList(0)
	for i := 0; i < testListLen; i++ {
		list.appendToTail(i)
	}

	return list
}
