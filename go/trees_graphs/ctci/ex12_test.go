package ctci

import "testing"

const (
	targetSum = 15
	want      = 5
)

func TestCountPathsWithSumBruteForce(t *testing.T) {
	root := getTestTree()
	got := CountPathsWithSumBruteForce(root, targetSum)
	if got != want {
		t.Errorf("want %d paths with sum %d, got %d", want, targetSum, got)
	}
}

func TestCountPathsWithSumAccumulator(t *testing.T) {
	root := getTestTree()
	got := CountPathsWithSumAccumulator(root, targetSum)
	if got != want {
		t.Errorf("want %d paths with sum %d, got %d", want, targetSum, got)
	}
}

func TestCountPathsWithSumOptimized(t *testing.T) {
	root := getTestTree()
	got := CountPathsWithSumAccumulator(root, targetSum)
	if got != want {
		t.Errorf("want %d paths with sum %d, got %d", want, targetSum, got)
	}
}

func getTestTree() *BinaryTreeNode {
	root := &BinaryTreeNode{data: 5}
	rootLeft := &BinaryTreeNode{data: 5}
	rootLeftLeft := &BinaryTreeNode{data: 5}
	rootLeftRight := &BinaryTreeNode{data: 10}
	rootRight := &BinaryTreeNode{data: 10}
	rootRightLeft := &BinaryTreeNode{data: 0}
	rootRightRight := &BinaryTreeNode{data: 20}
	rootRightRightLeft := &BinaryTreeNode{data: -20}

	root.left = rootLeft
	rootLeft.left = rootLeftLeft
	rootLeft.right = rootLeftRight
	root.right = rootRight
	rootRight.left = rootRightLeft
	rootRight.right = rootRightRight
	rootRightRight.left = rootRightRightLeft

	return root
}
