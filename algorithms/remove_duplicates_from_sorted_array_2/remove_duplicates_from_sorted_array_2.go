package main

import "fmt"

// second version of remove duplicate and modifies the same array but this time
// remove and left at most 2 unique elements
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	removeDuplicates(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d, ", nums[i])
	}
}
