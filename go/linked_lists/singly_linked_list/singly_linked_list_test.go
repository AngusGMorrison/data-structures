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
		t.Errorf("want *Node %+v, got %+v", want, got)
	}
}

func TestAppendToHead(t *testing.T) {
	tailData := 1
	list := New(tailData)
	wantData := 2
	list = list.AppendToHead(wantData)

	if list.Data != wantData {
		t.Errorf("want head to have data %d, got %d", wantData, list.Data)
	}
	if list.Next.Data != tailData {
		t.Errorf("want tail to have data %d, got %d", tailData, list.Next.Data)
	}
}

func TestAppendToTail(t *testing.T) {
	headData := 1
	list := New(headData)
	wantData := 2
	list.AppendToTail(wantData)

	if list.Next.Data != wantData {
		t.Errorf("want tail to have data %d, got %d", wantData, list.Next.Data)
	}
	if list.Data != headData {
		t.Errorf("want head to have data %d, got %d", headData, list.Data)
	}
}

func TestDelete(t *testing.T) {
	t.Run("deletes from head of list", func(t *testing.T) {
		list := nNodeList(3)
		list = list.Delete(0)
		wantHead := 1

		if list.Data != wantHead {
			t.Errorf("want head to have data %d, got %d", wantHead, list.Data)
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
	for i := int(n) - 2; i >= 0; i-- {
		list = list.AppendToHead(i)
	}
	return list
}
