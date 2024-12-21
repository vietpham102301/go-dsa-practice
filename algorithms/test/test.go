package main

import "fmt"

func modifySlice(s []int) {
	copy(s, []int{10, 20, 30})
}

func main() {
	original := []int{1, 2, 3}
	newSlice := make([]int, 3)
	copy(newSlice, original)
	modifySlice(newSlice)
	fmt.Println(original) // Output: [1, 2, 3]
	fmt.Println(newSlice)
}
