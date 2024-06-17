package main

import "fmt"

const StackSize = 3

type Stack struct {
	arr  [StackSize]int
	curr int
}

type SetOfStack struct {
	stacks []*Stack
	curr   int
}

func (s *SetOfStack) Push(item int) {
	if s.stacks == nil {
		newStack := Stack{}
		newStack.arr[newStack.curr] = item
		newStack.curr++
		s.stacks = []*Stack{&newStack}

		return
	}

	currentStack := s.stacks[s.curr]

	if currentStack.curr >= StackSize {
		s.curr++

		newStack := Stack{}
		newStack.arr[newStack.curr] = item
		newStack.curr++

		s.stacks = append(s.stacks, &newStack)

		return
	}

	currentStack.arr[currentStack.curr] = item
	currentStack.curr++
}

func (s *SetOfStack) Pop() (int, bool) {
	if s.stacks == nil || s.curr < 0 {
		return -1, false
	}

	currentStack := s.stacks[s.curr]

	if currentStack.curr == 0 {
		s.stacks = s.stacks[:s.curr]
		s.curr--
	}

	if s.curr < 0 {
		return -1, false
	}

	currentStack = s.stacks[s.curr]

	data := currentStack.arr[currentStack.curr-1]
	currentStack.curr--

	return data, true
}

func (s *SetOfStack) IsEmpty() bool {
	cond := s.stacks == nil || len(s.stacks) < 1 || (len(s.stacks) == 1 && s.stacks[s.curr].curr == 0)
	return cond
}

func main() {
	myStacks := SetOfStack{}
	myStacks.Push(1)
	myStacks.Push(2)
	myStacks.Push(3)

	myStacks.Push(4)
	myStacks.Push(5)
	myStacks.Push(6)

	myStacks.Push(7)
	myStacks.Push(8)
	myStacks.Push(9)

	for !myStacks.IsEmpty() {
		item, _ := myStacks.Pop()
		fmt.Println(item)
	}
}
