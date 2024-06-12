package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
	data string
}

func (l *LinkedList) prepend(n *Node) {
	n.next = l.head
	if l.head != nil {
		l.head.prev = n
		if l.tail == nil {
			l.tail = l.head
		}
	}
	l.head = n
}

func (l LinkedList) printList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%s, ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func checkPalindrome(l LinkedList) bool {
	curr1, curr2 := l.head, l.tail

	for curr1 != nil && curr2 != nil {
		if curr1.data != curr2.data {
			return false
		}
		curr1 = curr1.next
		curr2 = curr2.prev
	}

	return true
}

func main() {
	node1 := &Node{data: "a"}
	node2 := &Node{data: "b"}
	node3 := &Node{data: "c"}
	node4 := &Node{data: "b"}
	node5 := &Node{data: "a"}

	linkedList := LinkedList{}
	linkedList.prepend(node5)
	linkedList.prepend(node4)
	linkedList.prepend(node3)
	linkedList.prepend(node2)
	linkedList.prepend(node1)

	linkedList.printList()

	fmt.Println(checkPalindrome(linkedList))
}
