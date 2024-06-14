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

func detectLoopNode(l LinkedList) *Node {
	curr := l.head
	visitedMap := make(map[*Node]bool)

	for curr != nil {
		isVisited := visitedMap[curr]
		if !isVisited {
			visitedMap[curr] = true
		} else {
			return curr
		}
		curr = curr.next
	}

	return nil
}

func detectLoopNode2(l LinkedList) *Node {
	slow, fast := l.head, l.head

	//detect cycle
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break
		}
	}
	if fast == nil || fast.next == nil {
		return nil
	}
	// find first point
	slow = l.head

	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow
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

	node7.next = node5
	//linkedList.printList()

	fmt.Println(detectLoopNode2(linkedList))
}
