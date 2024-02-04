package main

import (
	"fmt"
	"sort"
)

func removeDuplicatesIntuition(nums []int) int {
	frequencyMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := frequencyMap[nums[i]]
		if !ok {
			frequencyMap[nums[i]] = 1
		} else {
			frequencyMap[nums[i]]++
		}
	}

	index := 0
	for key := range frequencyMap {
		nums[index] = key
		index++
	}

	subset := nums[:index]
	sort.Ints(subset)
	copy(nums[:index], subset)

	return index
}

func removeDuplicatesOptimize(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicatesOptimize(nums)
	fmt.Println(nums)
}
