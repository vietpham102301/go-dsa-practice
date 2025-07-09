package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func linkedListReversal(head *Node) *Node {
	var prevNode *Node = nil
	currNode := head

	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode

		prevNode = currNode
		currNode = nextNode
	}

	return prevNode
}

func linkedListReversalRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := linkedListReversalRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func main() {
	head := &Node{Value: 1, Next: nil}
	head.Next = &Node{Value: 2, Next: nil}
	head.Next.Next = &Node{Value: 3, Next: nil}
	head.Next.Next.Next = &Node{Value: 4, Next: nil}

	newHead := linkedListReversal(head)
	currNode := newHead
	for currNode != nil {
		fmt.Printf("%v, ", currNode.Value)
		currNode = currNode.Next
	}
	fmt.Println()
	newHead = linkedListReversalRecursive(newHead)
	currNode = newHead
	for currNode != nil {
		fmt.Printf("%v, ", currNode.Value)
		currNode = currNode.Next
	}
}
