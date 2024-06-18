package main

import "fmt"

type Node struct {
	Next *Node
	Data int
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(item int) {
	newNode := &Node{Data: item, Next: s.head}
	s.head = newNode
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return -1, false
	}
	data := s.head.Data
	s.head = s.head.Next
	return data, true
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

type MyQueue struct {
	Stack1 Stack
	Stack2 Stack
}

func (q *MyQueue) Add(item int) {
	for !q.Stack2.IsEmpty() {
		data, s := q.Stack2.Pop()
		if s {
			q.Stack1.Push(data)
		}
	}

	q.Stack1.Push(item)
}

func (q *MyQueue) Remove() int {
	for !q.Stack1.IsEmpty() {
		data, s := q.Stack1.Pop()
		if s {
			q.Stack2.Push(data)
		}
	}

	res, _ := q.Stack2.Pop()

	return res
}

func main() {
	myQueue := MyQueue{}

	myQueue.Add(1)
	myQueue.Add(2)
	myQueue.Add(3)

	fmt.Println(myQueue.Remove())
	fmt.Println(myQueue.Remove())
	fmt.Println(myQueue.Remove())
	fmt.Println(myQueue.Remove())

	myQueue.Add(4)
	myQueue.Add(5)
	myQueue.Add(6)

	fmt.Println(myQueue.Remove())
	fmt.Println(myQueue.Remove())
	fmt.Println(myQueue.Remove())
}
