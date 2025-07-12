package main

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
	}
}

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

func (l *LRUCache) addToTail(node *Node) {
	node.prev = l.tail
	node.next = nil
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	if l.head == nil {
		l.head = node
	}
}

func (l *LRUCache) moveToTail(node *Node) {
	if node == l.tail {
		return
	}
	l.removeNode(node)
	l.addToTail(node)
}

func (l *LRUCache) Put(key, value int) {
	if node, exists := l.cache[key]; exists {
		node.value = value
		l.moveToTail(node)
		return
	}

	if len(l.cache) == l.capacity {
		delete(l.cache, l.head.key)
		l.removeNode(l.head)
	}

	newNode := &Node{key: key, value: value}
	l.addToTail(newNode)
	l.cache[key] = newNode
}

func (l *LRUCache) Get(key int) int {
	if node, exists := l.cache[key]; exists {
		l.moveToTail(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) printList() {
	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Printf("key: %d, val: %d ", curr.key, curr.value)
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

	cache.Get(3)
	cache.printList()

	res := cache.Get(10)
	fmt.Println(res)
}
