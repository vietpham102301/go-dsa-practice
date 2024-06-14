package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
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

func (l LinkedList) printList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d, ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}

func checkPalindrome(l LinkedList) bool {
	curr1, curr2 := l.head, l.tail

	for curr1 != nil && curr2 != nil {
		if curr1.data != curr2.data {
			return false
		}
		curr1 = curr1.next
		curr2 = curr2.prev
	}

	return true
}

type Stack []int

func (s *Stack) Push(item int) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func checkPalindrome2(l LinkedList) bool {
	s := Stack{}

	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		s.Push(slow.data)
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		top, _ := s.Pop()
		if slow.data != top {
			return false
		}
		slow = slow.next
	}

	return true
}

type Result struct {
	node   *Node
	result bool
}

// Function to check if a linked list is a palindrome
func isPalindrome(head *Node) bool {
	length := lengthOfList(head)
	res := isPalindromeRecurse(head, length)
	return res.result
}

// recursive approach
func isPalindromeRecurse(head *Node, length int) Result {
	if head == nil || length <= 0 { // Even number of nodes
		return Result{head, true}
	} else if length == 1 { // Odd number of nodes
		return Result{head.next, true}
	}

	// Recurse on sublist
	res := isPalindromeRecurse(head.next, length-2)

	// If child calls are not a palindrome, pass back up a failure
	if !res.result || res.node == nil {
		return res
	}

	// Check if matches corresponding node on other side
	if res.node != nil {
		res.result = head.data == res.node.data
		res.node = res.node.next
	}

	return res
}

// Function to calculate the length of the linked list
func lengthOfList(head *Node) int {
	size := 0
	for head != nil {
		size++
		head = head.next
	}
	return size
}

func main() {
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	//node3 := &Node{data: 3}
	node4 := &Node{data: 2}
	node5 := &Node{data: 1}

	linkedList := LinkedList{}
	linkedList.prepend(node5)
	linkedList.prepend(node4)
	//linkedList.prepend(node3)
	linkedList.prepend(node2)
	linkedList.prepend(node1)

	linkedList.printList()

	fmt.Println(isPalindrome(linkedList.head))
}
