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

func (l *LinkedList) deleteStartFromNode(prevNode *Node, n *Node, val int) {
	currentNode := n

	if currentNode.data == val {
		prevNode.next = currentNode.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.data == val {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

func (l LinkedList) printList() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d, ", currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) removeDups() {
	currentNode := l.head

	for currentNode != nil {
		l.deleteStartFromNode(currentNode, currentNode.next, currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) removeDups2() {
	hashMap := make(map[int]bool)

	curr := l.head

	hashMap[curr.data] = true

	for curr.next != nil {
		isAppeared := hashMap[curr.next.data]
		if isAppeared {
			curr.next = curr.next.next
			continue
		} else {
			hashMap[curr.next.data] = true
		}
		curr = curr.next
	}
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

	linkedList.removeDups2()
	fmt.Println()
	fmt.Println("after remove dups")
	linkedList.printList()
}
