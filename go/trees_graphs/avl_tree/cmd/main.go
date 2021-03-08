package main

import (
	avl "github.com/angusgmorrison/datastructs/trees_graphs/avl_tree"
)

func main() {
	// Insertion with left rotation
	var root *avl.Node
	root, _ = root.Insert(13)
	root, _ = root.Insert(10)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(14)
	root, _ = root.Insert(20)
	root, _ = root.Insert(18)
	root, _ = root.Insert(25)
	root, _ = root.Insert(30)
	root.Print(0)

	// Insertion with right rotation
	root = nil
	root, _ = root.Insert(13)
	root, _ = root.Insert(10)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(11)
	root, _ = root.Insert(16)
	root, _ = root.Insert(4)
	root, _ = root.Insert(8)
	root, _ = root.Insert(3)
	root.Print(0)

	// Insertion with left-right rotation
	root = nil
	root, _ = root.Insert(13)
	root, _ = root.Insert(10)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(14)
	root, _ = root.Insert(20)
	root, _ = root.Insert(18)
	root, _ = root.Insert(19)
	root.Print(0)

	// Insertion with right-left rotation
	root = nil
	root, _ = root.Insert(13)
	root, _ = root.Insert(8)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(12)
	root, _ = root.Insert(16)
	root, _ = root.Insert(7)
	root, _ = root.Insert(6)
	root.Print(0)

	// Deletion with left rotation
	root = nil
	root, _ = root.Insert(10)
	root, _ = root.Insert(8)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(14)
	root, _ = root.Insert(16)
	root, _ = root.Insert(20)
	root, _ = root.Delete(8)
	root.Print(0)

	// Deletion with right rotation
	root = nil
	root, _ = root.Insert(10)
	root, _ = root.Insert(7)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(8)
	root, _ = root.Insert(16)
	root, _ = root.Insert(1)
	root, _ = root.Delete(15)
	root.Print(0)

	// Deletion without rotation (two-child deletion)
	root = nil
	root, _ = root.Insert(10)
	root, _ = root.Insert(7)
	root, _ = root.Insert(15)
	root, _ = root.Insert(14)
	root, _ = root.Insert(16)
	root, _ = root.Delete(15)
	root.Print(0)

	// Deletion with left-right rotation
	root = nil
	root, _ = root.Insert(10)
	root, _ = root.Insert(7)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(8)
	root, _ = root.Insert(16)
	root, _ = root.Insert(9)
	root, _ = root.Delete(15)
	root.Print(0)

	// Deletion with right-left rotation
	root = nil
	root, _ = root.Insert(10)
	root, _ = root.Insert(8)
	root, _ = root.Insert(15)
	root, _ = root.Insert(5)
	root, _ = root.Insert(14)
	root, _ = root.Insert(16)
	root, _ = root.Insert(12)
	root, _ = root.Delete(8)
	root.Print(0)
}
