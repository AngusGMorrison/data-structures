// Package redblack implements a red-black tree, a balanced binary search tree
// that enables search, insertion and deletion in O(logn) time.
//
// Although less perfectly balanced than AVL trees, they require fewer rotations
// to remain in balance and thus are better suited to uses which involve more
// modifications than lookups.
//
// * Every node must be either red or black.
// * Red nodes cannot be adjacent to each other (a red node cannot have a red
//   parent or child).
// * Every path from a node to any of its descendant nil nodes has the same
//   number of black nodes (nil nodes are counted as black nodes).
package redblack
