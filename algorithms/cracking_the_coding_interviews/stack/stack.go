package stack

type Node[T any] struct {
	Next *Node[T]
	Data T
}

type Stack[T any] struct {
	Head *Node[T]
}

func (s *Stack[T]) Push(item T) {
	newNode := &Node[T]{Data: item}
	newNode.Next = s.Head

	s.Head = newNode
}

func (s *Stack[T]) Pop() (*T, bool) {
	top := s.Head
	if top == nil {
		return nil, false
	}

	s.Head = s.Head.Next

	return &top.Data, true
}

func (s Stack[T]) Peek() (*T, bool) {
	if s.Head == nil {
		return nil, false
	}

	return &s.Head.Data, true
}

func (s Stack[T]) IsEmpty() bool {
	return s.Head == nil
}

//func main() {
//	myStack := Stack[int]{}
//
//	myStack.Push(1)
//	myStack.Push(2)
//	myStack.Push(3)
//
//	for !myStack.IsEmpty() {
//		item, _ := myStack.Pop()
//		fmt.Println(*item)
//	}
//
//}
