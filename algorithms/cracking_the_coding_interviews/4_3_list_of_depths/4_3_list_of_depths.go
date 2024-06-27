package main

import (
	"fmt"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

type LinkedListNode struct {
	next *LinkedListNode
	data *TreeNode
}

type LinkedList struct {
	head *LinkedListNode
}

func (l *LinkedList) prepend(n *TreeNode) {
	newNode := &LinkedListNode{data: n}
	newNode.next = l.head
	l.head = newNode
}

func createLinkedListAtEachDepthRecursion(rootNode *TreeNode, lists *[]*LinkedList, level int) {
	if rootNode == nil {
		return
	}
	var list *LinkedList
	if level == len(*lists) {
		list = &LinkedList{}
		*lists = append(*lists, list)
	} else {
		list = (*lists)[level]
	}
	list.prepend(rootNode)
	createLinkedListAtEachDepthRecursion(rootNode.left, lists, level+1)
	createLinkedListAtEachDepthRecursion(rootNode.right, lists, level+1)
}

func createLinkedListBFS(rootNode *TreeNode) []LinkedList {
	var res []LinkedList

	currList := LinkedList{}
	if rootNode != nil {
		currList.prepend(rootNode)
	}

	for currList.head != nil {
		res = append(res, currList)
		parents := currList
		currList = LinkedList{}

		p := parents.head
		for p != nil {
			if p.data.left != nil {
				currList.prepend(p.data.left)
			}
			if p.data.right != nil {
				currList.prepend(p.data.right)
			}
			p = p.next
		}
	}

	return res
}

func printList(l LinkedList) {
	curr := l.head

	for curr != nil {
		fmt.Printf("%d ->", curr.data.data)
		curr = curr.next
	}
	fmt.Println()
}

func main() {
	root := &TreeNode{data: 4}
	root.left = &TreeNode{data: 2}
	root.right = &TreeNode{data: 5}
	root.left.left = &TreeNode{data: 1}
	root.left.right = &TreeNode{data: 3}

	// res := &[]*LinkedList{}

	res := createLinkedListBFS(root)

	for _, r := range res {
		printList(r)
	}
}
