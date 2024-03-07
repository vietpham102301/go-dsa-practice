package main

import (
	"fmt"
)

type MinHeap struct {
	arr []int
}

func (m *MinHeap) Insert(val int) {
	m.arr = append(m.arr, val)
	m.heapifyUp(len(m.arr) - 1)
}

func (m *MinHeap) Extract() int {

	element := m.arr[0]
	l := len(m.arr) - 1
	m.arr[0] = m.arr[l]
	m.arr = m.arr[:l]

	m.heapifyDown(0)

	return element
}

func (m *MinHeap) heapifyDown(index int) {
	if len(m.arr) == 0 {
		return
	}
	leftChild := getLeftChild(index)
	rightChild := getRightChild(index)

	var childToCompare int
	if leftChild == len(m.arr)-1 {
		childToCompare = leftChild
	} else if m.arr[leftChild] < m.arr[rightChild] {
		childToCompare = leftChild
	} else {
		childToCompare = rightChild
	}

	for m.arr[index] > m.arr[childToCompare] {
		m.swap(index, childToCompare)
		index = childToCompare
		leftChild = getLeftChild(index)
		rightChild = getRightChild(index)

		if leftChild > len(m.arr)-1 || rightChild > len(m.arr)-1 {
			return
		}

		if leftChild == len(m.arr)-1 {
			childToCompare = leftChild
		} else if m.arr[leftChild] < m.arr[rightChild] {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}
	}
}

func (m *MinHeap) heapifyUp(index int) {
	parent := getParent(index)
	for index > 0 {
		if m.arr[index] < m.arr[parent] {
			m.swap(index, parent)
			index = parent
			parent = getParent(index)
		} else {
			break
		}
	}
}

func (m *MinHeap) swap(i int, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func getParent(index int) int {
	return (index - 1) / 2
}

func getLeftChild(index int) int {
	return 2*index + 1
}

func getRightChild(index int) int {
	return 2*index + 2
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Printf("Error reading number of potions: %v\n", err)
		return
	}

	pq := MinHeap{}
	life := 0
	count := 0
	potions := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&potions[i]); err != nil {
			fmt.Printf("Error reading potion effect: %v\n", err)
			return
		}
		life += potions[i]
		count++
		pq.Insert(potions[i])
		if life < 0 {
			life -= pq.Extract()
			count--
		}
	}

	fmt.Println(count)
}
