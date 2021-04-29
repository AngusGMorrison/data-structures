package ctci

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateDepthListsRecursive(t *testing.T) {
	t.Run("nil tree", func(t *testing.T) {
		got := GenerateDepthListsRecursive(nil)
		if got != nil {
			t.Errorf("want nil, got %v", got)
		}
	})

	t.Run("non-nil tree", func(t *testing.T) {
		root := NewBinaryTreeNode(50, nil, nil)
		root.MassInsert(30, 10, 40, 70, 60, 80)

		wantNLists := 3
		got := GenerateDepthListsRecursive(root)
		if len(got) != wantNLists {
			t.Errorf("want %d lists, got %d:\n%s", wantNLists, len(got), depthListsToString(got))
		}

		depth1List := got[1]
		if depth1List.treeNode != root {
			t.Errorf("want list at depth 1 to contain the root node %+v, got %+v:\n%s",
				root, depth1List.treeNode, depth1List)
		}

		wantDepth3Elements := []*BinaryTreeNode{
			root.right.right,
			root.right.left,
			root.left.right,
			root.left.left,
		}

		depth3List := got[3]
		for cur, i := depth3List, 0; cur != nil; cur, i = cur.next, i+1 {
			if cur.treeNode != wantDepth3Elements[i] {
				t.Errorf("want list at depth 3 to contain node %+v at position %d, got %+v\n%s",
					wantDepth3Elements[i], i, cur, depth3List)
			}
		}
	})
}

func TestGenerateDepthListsBFS(t *testing.T) {
	t.Run("nil tree", func(t *testing.T) {
		got := GenerateDepthListsBFS(nil)
		if got != nil {
			t.Errorf("want nil, got %v", got)
		}
	})

	t.Run("non-nil tree", func(t *testing.T) {
		root := NewBinaryTreeNode(50, nil, nil)
		root.MassInsert(30, 10, 40, 70, 60, 80)

		wantNLists := 3
		got := GenerateDepthListsBFS(root)
		if len(got) != wantNLists {
			t.Errorf("want %d lists, got %d:\n%s", wantNLists, len(got), depthListsToString(got))
		}

		depth1List := got[1]
		if depth1List.treeNode != root {
			t.Errorf("want list at depth 1 to contain the root node %+v, got %+v:\n%s",
				root, depth1List.treeNode, depth1List)
		}

		wantDepth3Elements := []*BinaryTreeNode{
			root.right.right,
			root.right.left,
			root.left.right,
			root.left.left,
		}

		depth3List := got[3]
		for cur, i := depth3List, 0; cur != nil; cur, i = cur.next, i+1 {
			if cur.treeNode != wantDepth3Elements[i] {
				t.Errorf("want list at depth 3 to contain node %+v at position %d, got %+v\n%s",
					wantDepth3Elements[i], i, cur, depth3List)
			}
		}
	})
}

func depthListsToString(depthLists map[int]*BinaryTreeListNode) string {
	var sb strings.Builder
	for k, v := range depthLists {
		sb.WriteString(strconv.Itoa(k))
		sb.WriteString(": ")
		sb.WriteString(v.String())
		sb.WriteByte('\n')
	}

	return sb.String()
}
