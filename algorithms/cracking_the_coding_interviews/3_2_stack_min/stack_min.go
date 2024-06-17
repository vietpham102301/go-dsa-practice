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
	head    *Node
	minNode *Node
}

func InitStack() Stack {
	s := Stack{}
	s.minNode = &Node{data: math.MaxInt}

	return s
}

func (s *Stack) Push(item int) {
	newNode := &Node{data: item}

	newNode.next = s.head
	s.head = newNode

	if item <= s.minNode.data {
		newMinNode := &Node{data: item, next: s.minNode}
		s.minNode = newMinNode
	}
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return -1, false
	}

	data := s.head.data
	s.head = s.head.next

	if data == s.minNode.data {
		s.minNode = s.minNode.next
	}

	return data, true
}

func (s Stack) GetMinElement() (int, bool) {
	if s.minNode != nil {
		return s.minNode.data, true
	}
	return -1, false
}

func main() {
	s := InitStack()
	s.Push(100)
	s.Push(34)
	s.Push(12)
	s.Push(-199)
	s.Push(-2323)

	s.Pop()
	s.Pop()

	fmt.Println(s.GetMinElement())
}
