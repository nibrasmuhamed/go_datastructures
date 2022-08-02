package main

import "fmt"

type heap []int

func (h *heap) Init() *heap {
	*h = append(*h, 1, 2, 4)
	return h
}

func (h heap) Display() {
	for _, val := range (heap{}) {
		fmt.Printf("%d ->", val)
	}
}

func main() {
	myHeap := heap{}
	myHeap.Init()
	fmt.Println(myHeap)

}

// Min heap:
// root node, parent node has less value than child nodes.
//
// Max heap:
// highest value is in top
// Heap is represented as an array.
