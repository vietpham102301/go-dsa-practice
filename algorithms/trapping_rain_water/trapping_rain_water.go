package main

import (
	"fmt"
	"math"
)

func trap(height []int) int {
	res := 0
	cur := 0
	for cur < len(height) {
		itMayTrap := false
		minHeight := math.MaxInt
		minHeightIndex := -1
		nextBarIndex := -1

		if height[cur] > 0 {
			for j := cur + 1; j < len(height); j++ {
				if height[j] >= height[cur] {
					nextBarIndex = j
					break
				}
				if height[j] < height[cur] && minHeight > height[j] {
					itMayTrap = true
					minHeight = height[j]
					minHeightIndex = j
				}
			}
			if itMayTrap && nextBarIndex == -1 {
				nextBarHeight := height[minHeightIndex]
				for j := cur + 1; j < len(height); j++ {
					if height[j] < height[cur] && height[j] > nextBarHeight {
						nextBarHeight = height[j]
						nextBarIndex = j
					}
				}
			}
		} else {
			cur++
			continue
		}

		if itMayTrap && nextBarIndex != -1 {
			var area int
			var accBarArea int
			if height[cur] >= height[nextBarIndex] {
				area = height[nextBarIndex] * (nextBarIndex - cur + 1)
				accBarArea = height[nextBarIndex]
			} else {
				area = height[cur] * (nextBarIndex - cur + 1)
				accBarArea = height[cur]
			}

			for k := cur; k < nextBarIndex; k++ {
				if height[k] <= height[nextBarIndex] {
					accBarArea += height[k]
				} else {
					accBarArea += height[nextBarIndex]
				}
			}
			res += area - accBarArea
			cur = nextBarIndex
		} else {
			cur++
		}
	}
	return res
}

func main() {
	testCase := []int{4, 2, 3}

	fmt.Println(trap(testCase))
}
