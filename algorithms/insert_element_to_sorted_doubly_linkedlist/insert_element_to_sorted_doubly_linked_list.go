package main

import (
	"fmt"
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	if llist == nil {
		llist = &DoublyLinkedListNode{
			data: data,
			prev: nil,
			next: nil,
		}
		return llist
	}

	if llist.next == nil && llist.data < data {
		llist.next = &DoublyLinkedListNode{
			data: data,
			next: nil,
			prev: llist,
		}
		return llist
	} else if llist.next == nil && llist.data >= data {
		llist.prev = &DoublyLinkedListNode{
			data: data,
			next: llist,
			prev: nil,
		}
		return llist.prev
	}

	//travel the linked list to find the correct position to insert
	current := llist
	for current.next != nil && current.data < data {
		current = current.next
	}

	// if we insert at the head of the list current.prev = nil
	if current.prev == nil {
		current.prev = &DoublyLinkedListNode{
			data: data,
			next: current,
			prev: nil,
		}

		return current.prev

	}

	//insert at last position
	if current.next == nil && current.data < data {
		current.next = &DoublyLinkedListNode{
			data: data,
			next: nil,
			prev: current,
		}

		return llist
	}

	//insert at any position
	current.prev.next = &DoublyLinkedListNode{
		data: data,
		next: current,
		prev: current.prev,
	}

	current.prev = current.prev.next

	return llist

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
	llist = sortedInsert(llist, 1)
	llist = sortedInsert(llist, 2)
	llist = sortedInsert(llist, 3)
	llist = sortedInsert(llist, 4)

	// Print the list to verify the elements are inserted in sorted order
	printList(llist)
}
