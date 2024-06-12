package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	data int
}

func (l *LinkedList) prepend(n *Node) {
	n.next = l.head
	l.head = n
}

func (l LinkedList) printList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d, ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func (l *LinkedList) partition(val int) {
	if l.head == nil {
		return
	}

	var beforeStart, beforeEnd, afterStart, afterEnd *Node

	curr := l.head
	for curr != nil {
		next := curr.next
		curr.next = nil
		if curr.data < val {
			if beforeStart == nil {
				beforeStart = curr
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = curr
				beforeEnd = curr
			}
		} else {
			if afterStart == nil {
				afterStart = curr
				afterEnd = afterStart
			} else {
				afterEnd.next = curr
				afterEnd = curr
			}
		}
		curr = next
	}

	if beforeStart == nil {
		l.head = afterStart
	} else {
		beforeEnd.next = afterStart
		l.head = beforeStart
	}
}

func main() {
	node1 := &Node{data: 3}
	node2 := &Node{data: 5}
	node3 := &Node{data: 8}
	node4 := &Node{data: 5}
	node5 := &Node{data: 10}
	node6 := &Node{data: 2}
	node7 := &Node{data: 1}

	linkedList := LinkedList{}
	linkedList.prepend(node7)
	linkedList.prepend(node6)
	linkedList.prepend(node5)
	linkedList.prepend(node4)
	linkedList.prepend(node3)
	linkedList.prepend(node2)
	linkedList.prepend(node1)

	fmt.Println("Original list:")
	linkedList.printList()

	linkedList.partition(5)

	fmt.Println("Partitioned list:")
	linkedList.printList()
}
