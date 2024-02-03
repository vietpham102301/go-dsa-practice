package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(e int) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() int {
	l := len(s.elements)
	toRemove := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return toRemove
}

func main() {
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s)

	fmt.Println(s.Pop())
	s.Pop()
	//s.Pop()
	fmt.Println(s)
}
