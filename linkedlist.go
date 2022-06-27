package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) deleteNode(data int) {
	temp := l.head
	if temp.value == data {
		l.head = temp.next
		return
	}
	var prev *Node // = nil
	for temp.value != data && temp != nil {
		prev = temp
		temp = temp.next
	}

	if temp == nil {
		return
	}

	if temp == l.tail {
		l.tail = prev
		l.tail.next = nil
		return
	}

	prev.next = temp.next
}

func (l *List) addNode(data int) {
	newNode := new(Node)
	newNode.value = data
	temp := l.head
	if l.head == nil {
		l.head = newNode
		return
	}
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	l.tail = newNode
}

func (l List) display() {
	fmt.Println("Printing down the list")
	if l.head == nil {
		fmt.Println("The list is empty")
	} else {
		temp := l.head
		for temp != nil {
			fmt.Printf("%d -> ", temp.value)
			temp = temp.next
		}
	}
}

func (l *List) insertAfter(nextTo, data int) {
	newNode := new(Node)
	newNode.value = data
	temp := l.head
	for temp.value != nextTo && temp != nil {
		temp = temp.next
	}
	if temp == nil {
		// l.tail.next = newNode
		// l.tail = newNode
		return
	}

	newNode.next = temp.next
	temp.next = newNode
}

func (l *List) removeDuplicate() {
	var nextNode *Node
	temp := l.head
	for temp != nil {
		nextNode = temp.next
		for nextNode != nil && nextNode.value == temp.value {
			nextNode = nextNode.next
		}
		temp.next = nextNode
		if nextNode == l.tail {
			l.tail = temp
		}
		temp = nextNode
	}

}

func (l *List) reverseList() {
	var before, after *Node
	for l.head != nil {
		after = l.head.next
		l.head.next = before // before is null (making head's next nil)
		before = l.head
		l.head = after
	}
	l.head = before
}

func main() {
	mylist := List{}
	mylist.addNode(6)
	mylist.addNode(10)
	mylist.addNode(8)
	mylist.addNode(8)
	mylist.addNode(8)
	mylist.addNode(5)
	mylist.addNode(5)
	// mylist.removeDuplicate()
	mylist.reverseList()
	mylist.display()

}
