package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) deleteNode(data int) {
	temp := l.head
	if l.head.data == data {
		l.head = l.head.next
		l.head.prev = nil
		return
	} else {
		for temp != nil && temp.data != data {
			temp = temp.next
		}
	}
	if temp == l.tail {
		l.tail = temp.prev
		l.tail.next = nil
		temp.prev = nil
		return
	}
	temp.next.prev = temp.prev
	temp.prev.next = temp.next
}

func (l *List) addNode(data int) {
	//creating a new node
	temp := new(Node)
	temp.data = data
	// checking linked list is initialised or not.
	if l.head == nil {
		l.head = temp
	} else {
		l.tail.next = temp
		temp.prev = l.tail
	}
	l.tail = temp
}

func (l List) Display() {
	if l.head == nil {
		fmt.Println("Linked list is not yet initialised.")
	} else {
		temp := l.head
		for temp != nil {
			fmt.Printf("%d -> ", temp.data)
			temp = temp.next
		}
	}
}

func (l List) Reverse() {
	if l.head == nil {
		fmt.Println("Linked list is not yet initialised.")
	} else {
		temp := l.tail
		for temp != nil {
			fmt.Printf("%d -> ", temp.data)
			temp = temp.prev
		}
	}
}

func (l *List) insertAfter(nextTo, data int) {
	temp := l.head
	newNode := new(Node)
	newNode.data = data
	for temp != nil && temp.data != nextTo {
		temp = temp.next
	}
	if l.tail == temp {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		return
	}
	newNode.next = temp.next
	temp.next = newNode
	newNode.prev = temp
	newNode.next.prev = newNode
}

func (l *List) insertBefore(beforeTo, data int) {
	newNode := new(Node)
	newNode.data = data
	if l.head.data == beforeTo {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
		return
	}
	temp := l.head
	for temp != nil && temp.data != beforeTo {
		temp = temp.next
	}
	temp.prev.next = newNode
	newNode.prev = temp.prev
	temp.prev = newNode
	newNode.next = temp
}

func main() {
	list := List{}
	list.addNode(6)
	list.addNode(10)
	list.addNode(5)
	list.addNode(5)
	list.addNode(5)
	list.addNode(8)
	list.addNode(8)
	list.Display()
}
