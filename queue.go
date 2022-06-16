package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(data int) {
	temp := new(Node)
	temp.data = data
	if q.front == nil {
		q.front = temp
		return
	}
	curr := q.front
	for curr != nil {
		curr = curr.next
	}
	temp = curr.next
}

func (q *Queue) Display() {
	temp := q.front
	if temp == nil {
		fmt.Println("queue empty")
		return
	}
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(12)
	myQueue.Enqueue(12)

	myQueue.Display()

}
