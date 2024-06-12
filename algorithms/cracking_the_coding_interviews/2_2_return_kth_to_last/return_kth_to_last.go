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
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d, ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l LinkedList) returnKthToLast(k int) *LinkedList {
	result := &LinkedList{}

	curr := l.head
	count := 0
	for curr != nil {
		if count == k {
			result.head = curr
		}
		count++
		curr = curr.next
	}

	return result
}

func main() {
	node1 := &Node{data: 1}
	node2 := &Node{data: 1}
	node3 := &Node{data: 2}
	node4 := &Node{data: 2}
	node5 := &Node{data: 3}
	node6 := &Node{data: 3}

	linkedList := LinkedList{}

	linkedList.prepend(node1)
	linkedList.prepend(node6)
	linkedList.prepend(node5)
	linkedList.prepend(node3)
	linkedList.prepend(node4)
	linkedList.prepend(node2)

	linkedList.printList()

	kthToLast := linkedList.returnKthToLast(3)

	kthToLast.printList()
}
