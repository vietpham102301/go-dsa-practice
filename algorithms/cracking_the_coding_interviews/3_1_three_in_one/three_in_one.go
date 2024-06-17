package main

import "fmt"

type Stack struct {
	arr   [99]int
	curr1 int
	curr2 int
	curr3 int
}

func InitStack() Stack {
	s := Stack{}
	s.curr1 = 0
	s.curr2 = 33
	s.curr3 = 66

	return s
}

func (s *Stack) Push(item, part int) {
	if part == 1 && s.curr1 <= 32 {
		s.arr[s.curr1] = item
		s.curr1++
	} else if part == 2 && s.curr2 <= 65 {
		s.arr[s.curr2] = item
		s.curr2++
	} else if part == 3 && s.curr3 <= 98 {
		s.arr[s.curr3] = item
		s.curr3++
	} else {
		fmt.Printf("Stack %d is full or invalid stack part\n", part)
	}
}

func (s *Stack) Pop(part int) (int, bool) {
	if part == 1 && s.curr1 != 0 {
		data := s.arr[s.curr1-1]
		s.curr1--
		return data, true
	} else if part == 2 && s.curr2 != 33 {
		data := s.arr[s.curr2-1]
		s.curr2--
		return data, true
	} else if part == 3 && s.curr3 != 66 {
		data := s.arr[s.curr3-1]
		s.curr3--
		return data, true
	}

	return -1, false
}

func (s Stack) Peek(part int) (int, bool) {
	if part == 1 && s.curr1 != 0 {
		return s.arr[s.curr1-1], true
	} else if part == 2 && s.curr2 != 33 {
		return s.arr[s.curr2-1], true
	} else if part == 3 && s.curr3 != 66 {
		return s.arr[s.curr3-1], true
	}

	return -1, false
}

func (s Stack) IsEmpty(part int) bool {
	if part == 1 {
		return s.curr1 == 0
	} else if part == 2 {
		return s.curr2 == 33
	} else {
		return s.curr3 == 66
	}
}

func main() {
	s := InitStack()

	s.Push(1, 1)
	s.Push(2, 1)
	s.Push(3, 1)

	s.Push(1, 2)
	s.Push(2, 2)
	s.Push(3, 2)

	s.Push(1, 3)
	s.Push(2, 3)
	s.Push(3, 3)

	for !s.IsEmpty(1) {
		item, _ := s.Pop(1)
		fmt.Println(item)
	}
	for !s.IsEmpty(2) {
		item, _ := s.Pop(2)
		fmt.Println(item)
	}
	for !s.IsEmpty(3) {
		item, _ := s.Pop(3)
		fmt.Println(item)
	}
}
