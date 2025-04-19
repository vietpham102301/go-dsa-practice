package main

import (
	"fmt"
	"slices"
)

func tripletSum(arr []int, target int) [][]int {
	slices.Sort(arr)
	var res [][]int

	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue // skip duplicates
		}

		// Two pointer approach
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				res = append(res, []int{arr[i], arr[left], arr[right]})
				left++
				right--
				// skip duplicates
				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for left < right && arr[right] == arr[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	testCases := [][]int{
		{},
		{0},
		{1, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 0, 1, -1, 1, -1},
	}

	for _, testCase := range testCases {
		res := tripletSum(testCase, 0)
		fmt.Println(res)
	}
}
