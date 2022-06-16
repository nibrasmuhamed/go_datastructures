package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(data int) {
	temp := new(Node)
	temp.data = data
	if s.top == nil {
		s.top = temp
	} else {
		temp.next = s.top
		s.top = temp
	}
}

func (s Stack) Display() {
	temp := s.top
	if temp == nil {
		fmt.Println("stack is empty")
	} else {
		for temp != nil {
			fmt.Printf("%d -> ", temp.data)
			temp = temp.next
		}
	}
}

func (s *Stack) Pop() {
	if s.top == nil {
		fmt.Println("stack underflow")
	} else {
		s.top = s.top.next
	}
}

func main() {
	myStack := Stack{}
	myStack.Push(10)
	myStack.Push(11)
	myStack.Push(12)
	myStack.Push(13)
	myStack.Pop()
	myStack.Display()
}
