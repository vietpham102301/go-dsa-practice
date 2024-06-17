package main

import (
	"fmt"
	"github.com/vietpham102301/go-dsa-practice/algorithms/cracking_the_coding_interviews/queue"
)

func main() {
	myQueue := queue.Queue[int]{}

	myQueue.Add(1)
	myQueue.Add(2)
	myQueue.Add(3)

	for !myQueue.IsEmpty() {
		item, _ := myQueue.Remove()
		fmt.Println(*item)
	}
}
