package main

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity: cap,
		cache:    make(map[int]*Node),
	}
}

func (l *LRUCache) Put(key, val int) {
	if node, exists := l.cache[key]; exists {
		node.value = val
		l.moveToTail(node)
		return
	}

	if len(l.cache) == l.capacity {
		delete(l.cache, l.head.key)
		l.removeNode(l.head)
	}

	node := &Node{key: key, value: val}
	l.addToTail(node)
	l.cache[key] = node
}

func (l *LRUCache) Get(key int) int {
	if node, exists := l.cache[key]; exists {
		l.moveToTail(node)
		return node.value
	}

	return -1
}

// helper
func (l *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}

func (l *LRUCache) moveToTail(node *Node) {
	if l.tail == node {
		return
	}

	l.removeNode(node)
	l.addToTail(node)
}

func (l *LRUCache) addToTail(node *Node) {
	node.next = nil
	node.prev = l.tail
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node

	if l.head == nil {
		l.head = node
	}
}

func (l *LRUCache) printList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("key: %d - val: %d, ", curr.key, curr.value)
		curr = curr.next
	}

	fmt.Println()
}

func main() {
	cache := NewLRUCache(3)

	cache.Put(1, 200)
	cache.Put(2, 500)
	cache.Put(3, 600)
	cache.printList()

	cache.Put(4, 700)
	cache.printList()

	cache.Put(2, 800)
	cache.printList()

	res := cache.Get(3)
	fmt.Println(res)
	cache.printList()

	res = cache.Get(10)
	fmt.Println(res)
}
