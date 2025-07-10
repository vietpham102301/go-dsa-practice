package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func linkedListIntersection(headA, headB *Node) *Node {
	ptrA, ptrB := headA, headB

	for ptrA != ptrB {
		ptrA = nextOrHead(ptrA, headB)
		ptrB = nextOrHead(ptrB, headA)
	}

	return ptrA
}

func nextOrHead(node, head *Node) *Node {
	if node != nil {
		return node.Next
	}
	return head
}

func main() {
	headA := &Node{Value: 1, Next: nil}
	headA.Next = &Node{Value: 2, Next: nil}
	headA.Next.Next = &Node{Value: 3, Next: nil}
	headA.Next.Next.Next = &Node{Value: 4, Next: nil}
	headA.Next.Next.Next.Next = &Node{Value: 5, Next: nil}
	headA.Next.Next.Next.Next.Next = &Node{Value: 6, Next: nil}
	headA.Next.Next.Next.Next.Next.Next = &Node{Value: 7, Next: nil}

	headB := &Node{Value: 1, Next: nil}
	headB.Next = &Node{Value: 2, Next: nil}
	headB.Next.Next = headA.Next.Next.Next.Next.Next

	result := linkedListIntersection(headA, headB)
	fmt.Println(result.Value)

}
