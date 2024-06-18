package main

import (
	"fmt"
	"math"
)

type Node struct {
	next *Node
	data int
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(item int) {
	newNode := &Node{data: item, next: s.head}
	s.head = newNode
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return -1, false
	}

	data := s.head.data
	s.head = s.head.next

	return data, true
}

func (s Stack) Peek() (int, bool) {
	if s.head == nil {
		return -1, false
	}

	return s.head.data, true
}

func (s Stack) IsEmpty() bool {
	return s.head == nil
}

func sortStack(s Stack, res *Stack) Stack {
	temp1 := Stack{}
	if s.IsEmpty() {
		return *res
	}
	maxValue := math.MinInt
	for !s.IsEmpty() {
		value, success := s.Pop()
		if success {
			if value >= maxValue {
				maxValue = value
			}
			temp1.Push(value)
		}
	}

	for !temp1.IsEmpty() {
		data, success2 := temp1.Pop()
		if success2 && data != maxValue {
			s.Push(data)
		}
	}

	res.Push(maxValue)
	sortStack(s, res)

	return *res
}

func main() {
	s := Stack{}
	s.Push(4)
	s.Push(5)
	s.Push(2)
	s.Push(3)
	s.Push(1)

	res := Stack{}
	res = sortStack(s, &res)

	for !res.IsEmpty() {
		data, _ := res.Pop()
		fmt.Println(data)
	}

}
