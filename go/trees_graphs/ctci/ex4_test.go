package ctci

import "testing"

func TestCheckBalanced(t *testing.T) {
	t.Run("tree is nil", func(t *testing.T) {
		got := CheckBalanced(nil)
		if got != true {
			t.Errorf("want true, got false")
		}
	})

	t.Run("tree has single element", func(t *testing.T) {
		root := NewBinaryTreeNode(1, nil, nil)
		got := CheckBalanced(root)
		if got != true {
			t.Errorf("want true, got false")
		}
	})

	t.Run("imbalance <= 1", func(t *testing.T) {
		root := NewBinaryTreeNode(5, nil, nil)
		root.Insert(1)
		got := CheckBalanced(root)
		if got != true {
			t.Errorf("want true, got false")
		}
	})

	t.Run("imbalance > 1", func(t *testing.T) {
		root := NewBinaryTreeNode(50, nil, nil)
		root.MassInsert(40, 30, 42, 45, 60)
		got := CheckBalanced(root)
		if got != false {
			t.Errorf("want false, got true")
		}
	})
}
