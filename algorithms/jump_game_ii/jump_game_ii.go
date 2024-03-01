package main

import "fmt"

//func jump(nums []int) int {
//	goal := len(nums) - 1
//	count := 0
//	for goal != 0 {
//		goal = findMaxStepIndex(nums, goal)
//		count++
//	}
//
//	return count
//}
//
//func findMaxStepIndex(nums []int, goal int) int {
//	maxStep := 0
//	indexFound := -1
//	for i := len(nums) - 2; i >= 0; i-- {
//		if i+nums[i] >= goal && (goal-i) > maxStep {
//			maxStep = goal - i
//			indexFound = i
//		}
//	}
//
//	return indexFound
//}

func jump(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i]+i, nums[i-1])
	}

	count := 0
	idx := 0
	for idx < len(nums)-1 {
		count++
		idx = nums[idx]
	}

	return count
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testArray := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(testArray))
}
