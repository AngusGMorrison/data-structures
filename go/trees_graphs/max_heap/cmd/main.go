package main

import (
	"fmt"

	maxheap "github.com/angusgmorrison/datastructs/trees_graphs/max_heap"
)

func main() {
	heap := maxheap.New(0)
	err := heap.Insert(10)
	if err != nil {
		fmt.Println(err)
	}

	heap.Insert(4)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(2)
	heap.Insert(9)
	heap.Insert(18)
	heap.Print()
	fmt.Println()
	fmt.Println()

	heap.Pop()
	heap.Print()
}
