package main

import (
	"fmt"
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

// sortedInsert this is the incorrect solution, will fix that later
func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	node := &DoublyLinkedListNode{
		data: data,
		next: nil,
		prev: nil,
	}
	if llist == nil {
		llist = node
		return llist
	}

	currentElement := llist

	if currentElement.next == nil {
		if data < currentElement.data {
			currentElement.prev = node
			node.next = currentElement
			return node
		} else {
			currentElement.next = node
			node.prev = currentElement
			return currentElement
		}
	}
	head := currentElement

	for currentElement.next != nil {
		if data >= currentElement.data {
			currentElement = currentElement.next
		} else {
			isFirst := false
			node.next = currentElement
			if currentElement.prev != nil {
				currentElement.prev.next = node
				isFirst = true
			} else {
				return node
			}
			tempPrev := currentElement.prev
			currentElement.prev = node
			node.prev = tempPrev
			if isFirst == true {
				return node
			}
			return head
		}
	}
	if data > currentElement.data {
		currentElement.next = node
		node.prev = currentElement
	} else {
		currentElement.prev.next = node
		node.prev = currentElement.prev
		currentElement.prev = node
		node.next = currentElement

	}

	return head
}

func printList(llist *DoublyLinkedListNode) {
	current := llist
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	var llist *DoublyLinkedListNode

	// Insert some elements into the list
	llist = sortedInsert(llist, 2)
	llist = sortedInsert(llist, 3)
	llist = sortedInsert(llist, 4)
	llist = sortedInsert(llist, 1)

	// Print the list to verify the elements are inserted in sorted order
	printList(llist)
}
