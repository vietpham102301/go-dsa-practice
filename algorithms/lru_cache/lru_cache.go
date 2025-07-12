package main

import "fmt"

type Node struct {
	key, value int
	next, prev *Node
}

type LRUCache struct {
	capacity   int
	cacheMap   map[int]*Node
	head, tail *Node
}

func (l *LRUCache) printList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("key: %d, val: %d ", curr.key, curr.value)
		curr = curr.next
	}
}

func (l *LRUCache) moveToTail() {
	//todo: implement this
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]*Node),
	}
}

func (l *LRUCache) Put(key, value int) {
	if l.cacheMap[key] == nil && l.capacity > 0 {
		l.capacity--
		if l.head == nil || l.tail == nil {
			l.head = &Node{key: key, value: value, next: nil, prev: nil}
			l.tail = l.head

		} else {
			l.tail.next = &Node{key, value, nil, l.tail}
			l.tail = l.tail.next
		}
		l.cacheMap[key] = l.tail
		return
	} else if l.cacheMap[key] == nil && l.capacity == 0 {
		l.head.next.prev = nil
		l.head = l.head.next

		l.tail.next = &Node{key: key, value: value, next: nil, prev: l.tail}
		l.tail = l.tail.next
		l.cacheMap[key] = l.tail
		return
	}

	refNode := l.cacheMap[key]
	if refNode != nil {
		refNode.value = value
		if refNode == l.tail {
			return
		}

		if refNode == l.head {
			l.head = l.head.next
			l.head.prev = nil
		}

		l.tail.next = refNode
		refNode.prev = l.tail
		refNode.next = nil
		l.tail = l.tail.next
	}
}

func (l *LRUCache) Get(key int) int {
	refNode := l.cacheMap[key]
	if refNode == nil {
		return -1
	}

	if refNode == l.tail {
		return refNode.value
	}

	if refNode == l.head {
		l.head = l.head.next
		l.head.prev = nil
	}

	l.tail.next = refNode
	refNode.prev = l.tail
	refNode.next = nil
	l.tail = l.tail.next

	return refNode.value
}

func main() {
	cache := NewLRUCache(3)

	// add when not reached capacity
	cache.Put(1, 200)
	cache.Put(2, 500)
	cache.Put(3, 600)
	cache.printList()

	fmt.Println()

	//evict when reach the capacity no duplicate key
	cache.Put(4, 700)
	cache.printList()

	fmt.Println()

	//update the recently used list when put duplicate key
	cache.Put(2, 800)
	cache.printList()

	fmt.Println()
	//get the key: 3
	cache.Get(3)
	cache.printList()

	fmt.Println()

	//get the key doesn't contain in the cache
	res := cache.Get(10)
	fmt.Println(res)

}
