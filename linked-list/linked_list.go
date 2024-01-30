package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	n.next = l.head
	l.head = n
	l.length++
}

func (l LinkedList) printList() {
	element := l.head
	for l.length > 0 {
		fmt.Printf("%d ", element.data)
		element = element.next
		l.length--
	}

	fmt.Println()
}

func (l *LinkedList) delete(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	prevNodeToDelete := l.head
	for prevNodeToDelete.next.data != value {
		if prevNodeToDelete.next.next == nil {
			return
		}
		prevNodeToDelete = prevNodeToDelete.next
	}

	prevNodeToDelete.next = prevNodeToDelete.next.next
	l.length--
}

func main() {
	linkedList := LinkedList{}
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	linkedList.prepend(node1)
	linkedList.prepend(node2)
	linkedList.prepend(node3)

	fmt.Println(linkedList)
	fmt.Println("Elements: ")
	linkedList.printList()
	linkedList.delete(1)
	linkedList.printList()

	linkedList.delete(23)
	linkedList.printList()
	linkedList.delete(3)
	linkedList.printList()
	linkedList.delete(2)
	linkedList.printList()
	fmt.Println(linkedList)
	linkedList.delete(2323)
	linkedList.printList()
}
