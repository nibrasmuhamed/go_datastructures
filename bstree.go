package main

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
		temp = currentNode
		return
	}
	for true {
		if value < currentNode.data {
			if currentNode.left == nil {
				temp = currentNode.left
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				temp = currentNode.right
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

func main() {
	tree := BinaryTree{}
	tree.Insert(39)
}
