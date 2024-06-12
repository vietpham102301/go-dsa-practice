package main

import "fmt"

type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	next *Node
	data int
}

func (l *LinkedList) prepend(n *Node) {
	if n.next == nil {
		n.next = l.head
	} else {
		curr := n
		for curr.next != nil {
			l.length++
			curr = curr.next
		}
	}
	l.head = n
	l.length++
}

func (l LinkedList) printList() {
	curr := l.head

	for curr != nil {
		fmt.Printf("%d, ", curr.data)
		curr = curr.next
	}

	fmt.Println()
}

func findIntersectNode(l1, l2 LinkedList) *Node {
	curr1, curr2 := l1.head, l2.head

	if l1.length > l2.length {
		diff := l1.length - l2.length
		for i := 0; i < diff; i++ {
			curr1 = curr1.next
		}
	} else {
		diff := l2.length - l1.length
		for i := 0; i < diff; i++ {
			curr2 = curr2.next
		}
	}
	for curr1 != nil && curr2 != nil {
		if curr1 == curr2 {
			return curr1
		}
		curr1 = curr1.next
		curr2 = curr2.next
	}
	return nil
}

func main() {
	node1 := &Node{data: 1}
	node2 := &Node{data: 7}
	node3 := &Node{data: 6}

	node4 := &Node{data: 5}
	node5 := &Node{data: 2}
	node6 := &Node{data: 2}
	node7 := &Node{data: 3}

	myList1 := LinkedList{}
	myList2 := LinkedList{}

	myList1.prepend(node5)
	myList1.prepend(node4)
	myList1.prepend(node3)
	myList1.prepend(node2)
	myList1.prepend(node1)
	myList1.printList()

	myList2.prepend(node3)
	myList2.prepend(node7)
	myList2.prepend(node6)

	myList2.printList()

	res := findIntersectNode(myList1, myList2)
	fmt.Println(res)
}
