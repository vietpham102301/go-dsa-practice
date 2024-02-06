package main

import "fmt"

func majorityElement(nums []int) int {
	candidate, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}

		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func majorityElementBoyerMoore(nums []int) int {
	frequencyMap := map[int]float64{}
	for i := 0; i < len(nums); i++ {
		_, ok := frequencyMap[nums[i]]
		if !ok {
			frequencyMap[nums[i]] = 1
		} else {
			frequencyMap[nums[i]]++
		}
	}
	maxValue := -1.0
	maxKey := -1
	for key := range frequencyMap {
		if (frequencyMap[key] / 2.0) > maxValue {
			maxValue = frequencyMap[key] / 2.0
			maxKey = key
		}
	}
	return maxKey
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 2, 4, 4, 2, 2}
	fmt.Print(majorityElement(nums))
}
