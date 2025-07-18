package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func removeKthLastNode(head *Node, k int) *Node {
	dummyNode := &Node{Value: -1, Next: head}

	trailer := dummyNode
	leader := dummyNode
	for i := 0; i < k; i++ {
		leader = leader.Next
		if leader == nil {
			return head
		}
	}

	for leader.Next != nil {
		trailer = trailer.Next
		leader = leader.Next
	}

	// detach the trailer from the kth node from last
	trailer.Next = trailer.Next.Next

	return dummyNode.Next
}

func main() {
	head := &Node{Value: 1, Next: nil}
	head.Next = &Node{Value: 2, Next: nil}
	head.Next.Next = &Node{Value: 3, Next: nil}
	head.Next.Next.Next = &Node{Value: 4, Next: nil}
	head.Next.Next.Next.Next = &Node{Value: 5, Next: nil}

	head = removeKthLastNode(head, 3)
	currNode := head
	for currNode != nil {
		fmt.Printf("%v, ", currNode.Value)
		currNode = currNode.Next
	}
	fmt.Println()
}
