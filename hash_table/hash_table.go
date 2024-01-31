package main

import "fmt"

const ArraySize = 7

// HashTable Struct
type HashTable struct {
	arr [ArraySize]*Bucket
}

// Bucket Struct
type Bucket struct {
	head   *BucketNode
	length int
}

// BucketNode Struct
type BucketNode struct {
	key  string
	next *BucketNode
}

// InitializeHashTable initialize a HashTable
func InitializeHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.arr {
		result.arr[i] = &Bucket{}
	}
	return result
}

// hash
func hash(key string) int {
	var sum int
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Insert a key to a hashtable
func (h *HashTable) Insert(key string) {
	index := hash(key)
	//checking and handle collision
	if h.arr[index].head == nil {
		h.arr[index].head = &BucketNode{key: key}
		h.arr[index].length++
	} else {
		newNode := &BucketNode{key: key}
		newNode.next = h.arr[index].head
		h.arr[index].head = newNode
		h.arr[index].length++
	}
}

// Search for a bucketNode
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	currentNode := h.arr[index].head
	if currentNode != nil {
		for i := 0; i < h.arr[index].length; i++ {
			if currentNode.key == key {
				return true
			} else {
				currentNode = currentNode.next
			}
		}
	}

	return false
}

// Delete a bucketNode in the bucket
func (h *HashTable) Delete(key string) {
	index := hash(key)
	prevDeleteNode := h.arr[index].head
	if prevDeleteNode != nil && prevDeleteNode.key == key {
		h.arr[index].head = prevDeleteNode.next
		return
	}
	if prevDeleteNode.next != nil {
		for i := 0; i < h.arr[index].length; i++ {
			if prevDeleteNode.next.key == key {
				prevDeleteNode.next = prevDeleteNode.next.next
				h.arr[index].length--
				return
			}
		}
	}
}

func main() {
	hashTable := InitializeHashTable()

	fmt.Println(hashTable)
	hashTable.Insert("Viet")
	hashTable.Insert("Nam")

	fmt.Println(hashTable.Search("Viet"))
	hashTable.Delete("Viet")
	fmt.Println(hashTable.Search("Viet"))
	fmt.Println(hashTable.Search("Nam"))
	hashTable.Delete("Nam")
	fmt.Println(hashTable.Search("Nam"))

	fmt.Println(hashTable.Search("smt"))
	hashTable.Insert("smt")
	fmt.Println(hashTable.Search("smt"))

}
