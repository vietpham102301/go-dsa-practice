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

func (l *LinkedList) partitioner(val int) {
	curr := l.head
	flag := false
	var lastNode *Node
	for curr != nil {
		if curr.data >= val {
			newNode := &Node{data: curr.data}
			l.prepend(newNode)
			if !flag {
				lastNode = newNode
				flag = true
			}
		}
		curr = curr.next
	}
	if !flag {
		return
	}
	curr2 := lastNode.next

	for curr2 != nil {
		if curr2.data < val {
			newNode := &Node{data: curr2.data}
			l.prepend(newNode)
		}
		curr2 = curr2.next
	}

	lastNode.next = nil
}

func (l *LinkedList) partitioner2(val int) {
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

func (l *LinkedList) partitioner3(val int) {
	head := l.head
	tail := l.head

	curr := l.head

	for curr != nil {
		next := curr.next
		if curr.data < val {
			curr.next = head
			head = curr
		} else {
			tail.next = curr
			tail = curr
		}
		curr = next
	}

	tail.next = nil

	l.head = head
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

	linkedList.prepend(node1)
	linkedList.prepend(node6)
	linkedList.prepend(node5)
	linkedList.prepend(node3)
	linkedList.prepend(node4)
	linkedList.prepend(node2)
	linkedList.prepend(node7)

	linkedList.printList()
	linkedList.partitioner3(5)
	linkedList.printList()
}
