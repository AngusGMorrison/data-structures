package ctci

import (
	"container/list"
	"fmt"
	"sort"
	"testing"
)

func TestSequences(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
		want [][]int
	}{
		{
			desc: "root is nil",
			data: nil,
			want: [][]int{},
		},
		{
			desc: "root is only node",
			data: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			desc: "tree is balanced",
			data: []int{2, 1, 3},
			want: [][]int{
				{2, 1, 3},
				{2, 3, 1},
			},
		},
		{
			desc: "right-hook fail",
			data: []int{5, 1, 10, 15, 12},
			want: [][]int{
				{5, 1, 10, 15, 12},
				{5, 10, 1, 15, 12},
				{5, 10, 15, 1, 12},
				{5, 10, 15, 12, 1},
			},
		},
		{
			desc: "left-hook fail",
			data: []int{10, 5, 1, 12},
			want: [][]int{
				{10, 5, 1, 12},
				{10, 5, 12, 1},
				{10, 12, 5, 1},
			},
		},
		{
			desc: "tree is unbalanced",
			data: []int{2, 1, 3, 10, 7, 15},
			want: [][]int{
				{2, 1, 3, 10, 7, 15},
				{2, 1, 3, 10, 15, 7},
				{2, 3, 1, 10, 7, 15},
				{2, 3, 1, 10, 15, 7},
				{2, 3, 10, 1, 7, 15},
				{2, 3, 10, 1, 15, 7},
				{2, 3, 10, 7, 1, 15},
				{2, 3, 10, 7, 15, 1},
				{2, 3, 10, 15, 1, 7},
				{2, 3, 10, 15, 7, 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tree := buildTree(tc.data)
			gotSeqs := listsToSlices(tree.Sequences())
			sortPermutations(gotSeqs)
			fmt.Println(gotSeqs)

			for i, wantSeq := range tc.want {
				for j, wantElem := range wantSeq {
					if gotSeqs[i][j] != wantElem {
						t.Fatalf("want sequences %v, got %v", tc.want, gotSeqs)
					}
				}
			}
		})
	}
}

func buildTree(data []int) *BinaryTreeNode {
	var root *BinaryTreeNode
	if len(data) == 0 {
		root = nil
	} else {
		root = NewBinaryTreeNode(data[0], nil, nil, nil)
		if len(data) > 1 {
			root.MassInsert(data[1:]...)
		}
	}

	return root
}

func listsToSlices(lists []*list.List) [][]int {
	slices := [][]int{}

	for _, list := range lists {
		slice := make([]int, 0, list.Len())
		for cur := list.Front(); cur != nil; cur = cur.Next() {
			slice = append(slice, cur.Value.(int))
		}

		slices = append(slices, slice)
	}

	return slices
}

func sortPermutations(perms [][]int) {
	sort.Slice(perms, func(i, j int) bool {
		first := perms[i]
		second := perms[j]

		for k, data := range first {
			if data < second[k] {
				return true
			}

			if data > second[k] {
				return false
			}
		}

		return true
	})
}
