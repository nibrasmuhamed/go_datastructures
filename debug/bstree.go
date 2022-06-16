package main

import "fmt"

type BST struct {
	data  int
	left  *BST
	right *BST
}

type BinaryTree struct {
	root *BST
}

func (b *BinaryTree) Insert(value int) {
	currentNode := b.root
	temp := new(BST)
	temp.data = value
	if currentNode == nil {
		b.root = temp
		return
	}
	for {
		if value < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = temp
				return
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = temp
				return
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func (b *BinaryTree) Contains(value int) bool {
	temp := b.root
	for temp != nil {
		if temp.data > value {
			temp = temp.left
		} else if temp.data < value {
			temp = temp.right
		} else {
			return true
		}
	}
	return false
}

func main() {
	tree := BinaryTree{}
	tree.Insert(39)
	tree.Insert(12)
	tree.Insert(34)

	fmt.Println(tree.Contains(34))

}
