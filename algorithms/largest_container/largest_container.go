package main

import "fmt"

func largestContainer(heights []int) int {
	left, right := 0, len(heights)-1
	maxVolume := 0
	for left < right {
		v := minInt(heights[left], heights[right]) * (right - left)

		if v > maxVolume {
			maxVolume = v
		}

		if heights[left] > heights[right] {
			right--
		} else if heights[left] < heights[right] {
			left++
		} else {
			left++
			right--
		}
	}

	return maxVolume
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := [][]int{
		{},
		{1},
		{0, 1, 0},
		{3, 3, 3, 3},
		{1, 2, 3},
		{3, 2, 1},
	}

	for _, test := range testCases {
		res := largestContainer(test)
		fmt.Println(res)
	}
}
