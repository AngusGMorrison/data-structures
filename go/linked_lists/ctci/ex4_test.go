package ctci

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	impls := []struct {
		name string
		fn   func(*SNode, int) *SNode
	}{
		{"PartionStable", PartitionStable},
		{"PartionUnstable", PartitionUnstable},
	}

	input := newSingleFromList(3, 5, 8, 5, 10, 2, 1)
	partition := 5
	for _, impl := range impls {
		t.Run(fmt.Sprintf("%s(%v, %d)", impl.name, input, partition), func(t *testing.T) {
			want := []int{3, 1, 2, 10, 5, 5, 8}
			got := impl.fn(input, partition)
			if !input.assertOrder(want...) {
				t.Errorf("got order %s, want %v", got.String(), want)
			}
		})
	}
}
