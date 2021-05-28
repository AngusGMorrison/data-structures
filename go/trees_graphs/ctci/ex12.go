package ctci

// 4.12 Paths with Sum: You are given a binary tree in which each node contains
// an integer value (which might be positive or negative). Design an algorithm
// to count the number of paths that sum to a given value. The path does not
// need to start or end at the root or a leaf, but it must go downwards
// (traveraling only from parent nodes to child nodes).
//
// Assumptions:
//   * Nodes do not have links to their parents.
//   * The input is not necessarily a binary search tree.
//   * We may use additional data structures.
//   * A single node with the given value counts as a path.
//   * Nodes may have 0 value.

// CountPathsWithSumBruteForce looks at all possible paths by traversing to each
// node and recursively trying all paths downwards, tracking the sum of that
// path. Once the target sum is reached, the total is incremented.
//
// In a balanced tree, the length of each path in no more than log n, and the
// function is called on each of the n nodes. Therefore, time complexity is O(n
// logn). In the worst case (a tree consisting of a single branch), this rises
// to O(n^2).
//
// Space complexity is O(log n) for a balanced tree based on the recursive calls
// that must be held in memory. This may grow to O(n) for an unbalanced tree.
func CountPathsWithSumBruteForce(n *BinaryTreeNode, targetSum int) int {
	if n == nil {
		return 0
	}

	// Count paths with sum starting from the root.
	pathsFromRoot := countPathsWithSumFromNode(n, targetSum, 0)

	// Try the nodes on the left and right.
	pathsOnLeft := CountPathsWithSumBruteForce(n.left, targetSum)
	pathsOnRight := CountPathsWithSumBruteForce(n.right, targetSum)

	return pathsFromRoot + pathsOnLeft + pathsOnRight
}

func countPathsWithSumFromNode(n *BinaryTreeNode, targetSum int, currentSum int) int {
	if n == nil {
		return 0
	}

	currentSum += n.data

	var totalPaths int
	if currentSum == targetSum {
		totalPaths++
	}

	totalPaths += countPathsWithSumFromNode(n.left, targetSum, currentSum)
	totalPaths += countPathsWithSumFromNode(n.right, targetSum, currentSum)
	return totalPaths
}

// Like CountPathsWithSomeBruteForce, CountPathsWithSumAccumulator performs a
// logn operation for all n nodes in the tree. Instead of counting downwards,
// when it reaches a node, it sums all paths to that nodes from those that came
// before it.
func CountPathsWithSumAccumulator(n *BinaryTreeNode, sum int) int {
	var count int
	var countPathsWithSum func(n *BinaryTreeNode, accum []int)
	countPathsWithSum = func(n *BinaryTreeNode, accum []int) {
		if n == nil {
			return
		}

		accum = append(accum, n.data)
		var runningTotal int
		for i := len(accum) - 1; i >= 0; i-- {
			runningTotal += accum[i]
			if runningTotal == sum {
				count++
			}
		}

		countPathsWithSum(n.left, accum)
		countPathsWithSum(n.right, accum)
	}

	countPathsWithSum(n, []int{})
	return count
}

// CountPathsWithSumOptimized uses a map to keep track of the running total for
// each node in a path. Each node we encounter is added to the running total. We
// can determine the number of sub-paths with the desired total that end at the
// current node by looking up runningTotal - targetSum in the map. I.e. if any
// previous nodes in the path had a running total that is targetSum less than
// the current running total, the path from the previous node to the current
// node must sum to targetSum.
//
// Time complexity is O(n). Each node is visited once, doing O(1) work each
// time. Space complexity is O(log n) in a balanced tree, or up to O(n) for an
// unbalanced tree.
func CountPathsWithSumOptimized(n *BinaryTreeNode, targetSum int) int {
	pathCount := make(map[int]int)

	var countPaths func(n *BinaryTreeNode, runningSum int) int
	countPaths = func(n *BinaryTreeNode, runningSum int) int {
		if n == nil {
			return 0
		}

		runningSum += n.data

		// How many nodes are there in the current path whose sum is the
		// difference between the current total for the path and the total we're
		// looking for?
		totalPaths := pathCount[runningSum-targetSum]
		if runningSum == targetSum {
			totalPaths++
		}

		// Recording the current running total before recursing, and remove it
		// again once we're finished to prevent nodes in other paths from
		// incorrectly referring to it.
		pathCount[runningSum]++
		totalPaths += countPaths(n.left, runningSum)
		totalPaths += countPaths(n.right, runningSum)
		decrementOrDelete(pathCount, runningSum)

		return totalPaths
	}

	return countPaths(n, 0)
}

// decrementOrDelete allows us to keep space requirements to a maximum of O(log
// n) (for a balanced tree) by removing any running total that is no longer in
// use.
func decrementOrDelete(m map[int]int, k int) {
	if m[k] > 1 {
		m[k]--
	} else {
		delete(m, k)
	}
}
