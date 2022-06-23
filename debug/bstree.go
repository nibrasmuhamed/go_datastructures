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

func (b *BinaryTree) remove(data int) {
	b.RemoveHelper(data, b.root, nil)
}

func (b *BinaryTree) RemoveHelper(data int, currentNode *BST, parentNode *BST) {
	for currentNode != nil {
		if data < currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if data > currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			// condition 1 : current node has right and left child.
			if currentNode.right != nil && currentNode.left != nil {
				currentNode.data = getMinimumValue(currentNode.right)
				b.RemoveHelper(currentNode.data, currentNode.right, currentNode)
			} else { // condition 2 : current node has only one child. either left or right.
				if parentNode == nil { // c2.1 : if there is no parent node (node we are trying to remove is root node)
					if currentNode.right == nil { // c2.1.1 if there is no right child
						b.root = currentNode.left
					} else { // c2.1.2 if there is no left child
						b.root = currentNode.right
					}
				} else { // c 2.2 if parent node is present
					if parentNode.left == currentNode { // c2.2.1 where the current node is present (left to parent)
						if currentNode.right == nil { // c2.2.1.2 if right is nil
							parentNode.left = currentNode.left
						} else {
							parentNode.left = currentNode.right
						}
					} else { // c2.2.2 if left is empty
						if currentNode.right == nil {
							parentNode.right = currentNode.left
						} else {
							parentNode.right = currentNode.right
						}
					}
				}
			}
			break
		}
	}
}

// function to get minimun value node in binary search tree.
func getMinimumValue(currentNode *BST) int {
	// condition 1 : current node's left == nil
	if currentNode.left == nil {
		return currentNode.data
	} else {
		return getMinimumValue(currentNode.left)
	}
}

func (b BinaryTree) inOrder() {
	inOrderHelper(b.root)
}

func inOrderHelper(tempNode *BST) {
	if tempNode != nil {
		inOrderHelper(tempNode.left)
		fmt.Printf("%d ->", tempNode.data)
		inOrderHelper(tempNode.right)
	}
}

func main() {
	tree := BinaryTree{}
	tree.Insert(39)
	tree.Insert(12)
	tree.Insert(34)
	fmt.Println(tree.Contains(12))
	tree.remove(300)
	fmt.Println(tree.Contains(39))
	tree.inOrder()
}
