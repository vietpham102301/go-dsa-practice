package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	prev *Node
	next *Node
	data int
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

func (l *LinkedList) postpend(n *Node) {
	curr := l.head

	if curr == nil {
		n.next = l.head
		if l.head != nil {
			l.head.prev = n
			if l.tail == nil {
				l.tail = l.head
			}
		}
		l.head = n
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = n
	n.prev = curr
}

func (l LinkedList) printList() {
	curr := l.head

	for curr != nil {
		fmt.Printf("%d, ", curr.data)
		curr = curr.next
	}

	fmt.Println()
}

// other approach is to use recursion
func addTwoLinkedList1(l1, l2 LinkedList) LinkedList {
	curr1, curr2 := l1.tail, l2.tail

	base := 10

	num1 := 0
	for curr1 != nil {
		num1 += curr1.data
		if curr1.prev != nil {
			num1 *= base
		}
		curr1 = curr1.prev
	}

	num2 := 0
	for curr2 != nil {
		num2 += curr2.data
		if curr2.prev != nil {
			num2 *= base
		}
		curr2 = curr2.prev
	}

	sum := num1 + num2
	result := LinkedList{}
	for sum > 0 {
		rem := sum % base
		newNode := &Node{data: rem}
		result.postpend(newNode)
		sum /= base
	}

	return result

}

func addTwoLinkedList2(l1, l2 LinkedList) LinkedList {
	curr1, curr2 := l1.head, l2.head

	base := 10

	num1 := 0
	for curr1 != nil {
		num1 += curr1.data
		if curr1.next != nil {
			num1 *= base
		}
		curr1 = curr1.next
	}

	num2 := 0
	for curr2 != nil {
		num2 += curr2.data
		if curr2.next != nil {
			num2 *= base
		}
		curr2 = curr2.next
	}

	sum := num1 + num2
	result := LinkedList{}
	for sum > 0 {
		rem := sum % base
		newNode := &Node{data: rem}
		result.prepend(newNode)
		sum /= base
	}

	return result
}

func main() {
	node1 := &Node{data: 6}
	node2 := &Node{data: 1}
	node3 := &Node{data: 7}

	node4 := &Node{data: 2}
	node5 := &Node{data: 9}
	node6 := &Node{data: 5}

	myList1 := LinkedList{}
	myList2 := LinkedList{}

	myList1.prepend(node3)
	myList1.prepend(node2)
	myList1.prepend(node1)
	myList1.printList()

	myList2.prepend(node6)
	myList2.prepend(node5)
	myList2.prepend(node4)

	myList2.printList()

	result := addTwoLinkedList2(myList1, myList2)

	result.printList()
}
