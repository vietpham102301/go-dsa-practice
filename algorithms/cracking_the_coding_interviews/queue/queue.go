package queue

type Node[T any] struct {
	Next *Node[T]
	Data T
}

type Queue[T any] struct {
	First *Node[T]
	Last  *Node[T]
}

func (q *Queue[T]) Add(item T) {
	newNode := &Node[T]{Data: item}
	if q.Last != nil {
		q.Last.Next = newNode
	}

	q.Last = newNode
	if q.First == nil {
		q.First = q.Last
	}
}

func (q *Queue[T]) Remove() (*T, bool) {
	if q.First == nil {
		return nil, false
	}

	data := q.First.Data
	q.First = q.First.Next

	return &data, true
}

func (q Queue[T]) Peek() (*T, bool) {
	if q.First == nil {
		return nil, false
	}

	return &q.First.Data, true
}

func (q Queue[T]) IsEmpty() bool {
	return q.First == nil
}

//func main() {
//	myQueue := Queue[int]{}
//
//	myQueue.Add(1)
//	myQueue.Add(2)
//	myQueue.Add(3)
//
//	for !myQueue.IsEmpty() {
//		item, _ := myQueue.Remove()
//		fmt.Println(*item)
//	}
//}
