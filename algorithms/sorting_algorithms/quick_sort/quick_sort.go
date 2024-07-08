package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[high], arr[i] = arr[i], arr[high]

	return i
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)

		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func main() {
	test := []int{3, 2, 1, 4, 5}
	quickSort(test, 0, len(test)-1)
	for i := 0; i < len(test); i++ {
		fmt.Printf("%d, ", test[i])
	}
}
