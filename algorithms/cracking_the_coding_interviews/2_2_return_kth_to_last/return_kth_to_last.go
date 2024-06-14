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

// wrong problem intend
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

func returnKthToLastNode(l LinkedList, k int) *Node {
	curr := l.head
	if curr == nil {
		return nil
	}
	i := 0
	node := returnKthToLastRecursive(curr, k, &i)

	return node
}

func returnKthToLastRecursive(node *Node, k int, i *int) *Node {
	if node == nil {
		return nil
	}

	n := returnKthToLastRecursive(node.next, k, i)
	*i += 1
	if *i == k {
		return node
	}

	return n
}

func returnKthToLastIterative(l LinkedList, k int) *Node {
	curr1, curr2 := l.head, l.head

	for i := 0; i < k; i++ {
		curr1 = curr1.next
	}

	for curr1 != nil {
		curr2 = curr2.next
		curr1 = curr1.next
	}

	return curr2
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

	//kthToLast := linkedList.returnKthToLast(3)
	//
	//kthToLast.printList()
	//node := returnKthToLastNode(linkedList, 4)

	node := returnKthToLastIterative(linkedList, 4)
	fmt.Println(node.data)
}
