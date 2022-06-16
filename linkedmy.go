package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type List struct {
	head *node
	tail *node
}

func (l *List) addAtEnd(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
}

func (l *List) deletenode(n *node) {
	temp := l.head
	fmt.Println("first", *temp)
	if temp != nil && temp.data == n.data {
		l.head = temp.next
	}
	fmt.Println("between", temp)
	var prev node
	for temp != nil && temp.data != n.data {
		fmt.Println("inside", temp)
		prev = *temp
		fmt.Println("assign", temp)
		fmt.Println("prev", &prev)
		temp = temp.next
		fmt.Println("incri", temp)
	}
	fmt.Println("second", temp)
	fmt.Println("prev2", &prev)
	prev.next = temp.next
	fmt.Println("third", temp)
}

func (l List) display() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		temp := l.head
		for temp != nil {
			fmt.Printf("%d is element\n", temp.data)
			temp = temp.next
		}
	}
}

func main() {
	mylist := List{}
	node1 := &node{data: 39}
	node2 := &node{data: 34}
	node3 := &node{data: 35}
	mylist.display()
	mylist.addAtEnd(node1)
	mylist.addAtEnd(node2)
	mylist.addAtEnd(node3)
	mylist.addAtEnd(&node{data: 07})
	mylist.display()
	fmt.Println("---------")
	mylist.deletenode(node2)
	//mylist.deletenode(&node{data: 07})
	mylist.display()
}
