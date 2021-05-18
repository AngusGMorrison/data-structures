package ctci

import (
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
			gotSeqs := tree.Sequences()
			sortPermutations(gotSeqs)

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

func sortPermutations(perms [][]int) {
	sort.Slice(perms, func(i, j int) bool {
		first := perms[i]
		second := perms[j]

		for k, data := range first {
			if second[k] < data {
				return false
			}
		}

		return true
	})
}
