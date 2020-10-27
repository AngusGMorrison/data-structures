package linkedlist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	data := 1
	want := &Node{Data: data}
	got := New(data)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestAppendToTail(t *testing.T) {
	t.Run("when list has only one entry", func(t *testing.T) {
		wantData := 2
		list := New(1)
		list.AppendToTail(wantData)
		gotTail := list.Next

		if gotTail.Data != wantData {
			t.Errorf("want tail to have data %d, got %d", wantData, gotTail)
		}
		if gotTail.Next != nil {
			t.Errorf("want tail.Next == nil, got %+v", gotTail.Next)
		}
		if gotTail.Prev != list {
			t.Errorf("want tail to have previous node %+v, got %+v", gotTail.Prev, list)
		}
	})

	t.Run("when list has multiple entries", func(t *testing.T) {
		list := nNodeList(2)
		wantData := 2
		list.AppendToTail(wantData)
		gotTail := list.tail()
		wantPrevData := 1

		if gotTail.Data != wantData {
			t.Errorf("want tail to have data %d, got %d", wantData, gotTail.Data)
		}
		if gotTail.Next != nil {
			t.Errorf("want tail.Next == nil, got %+v", gotTail.Next)
		}
		if gotTail.Prev.Data != wantPrevData {
			t.Errorf("want tail to have previous node with data %d, got %d", wantPrevData, gotTail.Prev.Data)
		}

	})
}

func TestAppendToHead(t *testing.T) {
	headData := 1
	list := New(headData)
	wantData := 2
	gotList := list.AppendToHead(wantData)

	if gotList.Data != wantData {
		t.Errorf("want new head to have data %d, got %d", wantData, gotList.Data)
	}
	if gotList.Next != list {
		t.Errorf("wanted gotList.Next == %+v, got %+v", list, gotList.Next)
	}
}

func TestDelete(t *testing.T) {
	t.Run("deletes from head of list", func(t *testing.T) {
		list := nNodeList(3)
		list = list.Delete(0)
		wantHead := 1
		if list.Data != wantHead {
			t.Errorf("want list.Data == %d, got %d", wantHead, list.Data)
		}
		if list.Prev != nil {
			t.Errorf("want list.Prev == nil, got %+v", list.Prev)
		}
	})

	t.Run("deletes from tail of list", func(t *testing.T) {
		list := nNodeList(3)
		list = list.Delete(2)
		wantHead, wantNext := 0, 1
		if list.Data != wantHead {
			t.Errorf("want head to have data %d, got %d", wantHead, list.Data)
		}
		if list.Next.Data != wantNext {
			t.Errorf("want head.Next to have data %d, got %d", wantNext, list.Next.Data)
		}
		if deleted, err := list.GetIndex(2); err == nil {
			t.Errorf("want tail to be deleted, got data %d", deleted)
		}
	})

	t.Run("deletes from middle of list", func(t *testing.T) {
		list := nNodeList(3)
		list = list.Delete(1)
		wantHead, wantNext := 0, 2
		if list.Data != wantHead {
			t.Errorf("want head to have data %d, got %d", wantHead, list.Data)
		}
		if list.Next.Data != wantNext {
			t.Errorf("want head.Next to have data %d, got %d", wantNext, list.Next.Data)
		}
		if list.Next.Prev.Data != wantHead {
			t.Errorf("wanted head.Next.Prev to have data %d, got %d", wantHead, list.Next.Prev.Data)
		}
	})
}

func TestGetIndex(t *testing.T) {
	len := 3
	list := nNodeList(uint(len))
	for i := len - 1; i > 0; i-- {
		got, err := list.GetIndex(i)
		if err != nil {
			t.Errorf("want list.GetIndex(%d) to yield err nil, got %v", i, err)
		}
		if got != i {
			t.Errorf("want list.GetIndex(%d) == %d, got %d", i, i, got)
		}
	}
}

func nNodeList(n uint) *Node {
	list := New(int(n - 1))
	for i := int(n - 2); i >= 0; i-- {
		list = list.AppendToHead(i)
	}
	return list
}
