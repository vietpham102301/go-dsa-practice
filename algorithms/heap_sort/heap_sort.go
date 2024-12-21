package main

import "fmt"

// heapify chuyển đổi một mảng thành cây heap
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// heapSort sắp xếp mảng sử dụng thuật toán Heap Sort
func heapSort(arr []int) {
	n := len(arr)

	// Tạo cây heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Sắp xếp
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{9, 4, 3, 8, 10, 2, 5}
	fmt.Println("Mảng ban đầu:", arr)
	heapSort(arr)
	fmt.Println("Mảng sau khi sắp xếp:", arr)
}
