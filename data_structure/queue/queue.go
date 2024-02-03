package main

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(e int) {
	q.elements = append(q.elements, e)
}

func (q *Queue) Dequeue() int {
	toRemove := q.elements[0]
	q.elements = q.elements[1:]

	return toRemove
}

func main() {
	q := &Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q)

	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}
