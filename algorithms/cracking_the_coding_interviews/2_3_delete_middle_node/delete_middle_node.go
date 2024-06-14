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

// not the case problem intend
func (l *LinkedList) deleteNode(n *Node) {
	curr := l.head

	if l.head.data == n.data {
		l.head = l.head.next
	}

	for curr.next != nil {
		if curr.next.data == n.data {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

// can't delete if the node is the last node, if need to handle that maybe create a dummy node.
func deleteNode(n *Node) bool {
	if n == nil || n.next == nil {
		return false
	}

	nextNode := n.next
	n.data = nextNode.data
	n.next = nextNode.next

	return true
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

	midNode := node4

	//linkedList.deleteNode(midNode)
	deleteNode(midNode)
	linkedList.printList()
}
