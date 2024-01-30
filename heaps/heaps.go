package main

import (
	"fmt"
)

// MaxHeap struct has a slice that hold the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

}

// maxHeapify will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func leftChild(i int) int {
	return i*2 + 1
}

// get the right child index
func rightChild(i int) int {
	return i*2 + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("cannot extract from empty array")
		return -1
	}
	extracted := h.array[0]
	l := len(h.array) - 1

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(i int) {
	lastIndex := len(h.array) - 1
	l, r := leftChild(i), rightChild(i)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[childToCompare] > h.array[i] {
			h.swap(childToCompare, i)
			i = childToCompare
			l, r = leftChild(i), rightChild(i)
		} else {
			return
		}
	}
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 8, 112, 24}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 9; i++ {
		m.Extract()
		fmt.Println(m)
	}
	
}
