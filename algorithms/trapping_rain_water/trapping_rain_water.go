package main

import (
	"fmt"
)

// trap intuition
//func trap(height []int) int {
//	res := 0
//	cur := 0
//	for cur < len(height) {
//		itMayTrap := false
//		minHeight := math.MaxInt
//		minHeightIndex := -1
//		nextBarIndex := -1
//
//		if height[cur] > 0 {
//			for j := cur + 1; j < len(height); j++ {
//				if height[j] >= height[cur] {
//					nextBarIndex = j
//					break
//				}
//				if height[j] < height[cur] && minHeight > height[j] {
//					itMayTrap = true
//					minHeight = height[j]
//					minHeightIndex = j
//				}
//			}
//			if itMayTrap && nextBarIndex == -1 {
//				nextBarHeight := height[minHeightIndex]
//				for j := cur + 1; j < len(height); j++ {
//					if height[j] < height[cur] && height[j] > nextBarHeight {
//						nextBarHeight = height[j]
//						nextBarIndex = j
//					}
//				}
//			}
//		} else {
//			cur++
//			continue
//		}
//
//		if itMayTrap && nextBarIndex != -1 {
//			var area int
//			var accBarArea int
//			if height[cur] >= height[nextBarIndex] {
//				area = height[nextBarIndex] * (nextBarIndex - cur + 1)
//				accBarArea = height[nextBarIndex]
//			} else {
//				area = height[cur] * (nextBarIndex - cur + 1)
//				accBarArea = height[cur]
//			}
//
//			for k := cur; k < nextBarIndex; k++ {
//				if height[k] <= height[nextBarIndex] {
//					accBarArea += height[k]
//				} else {
//					accBarArea += height[nextBarIndex]
//				}
//			}
//			res += area - accBarArea
//			cur = nextBarIndex
//		} else {
//			cur++
//		}
//	}
//	return res
//}

// O(n) time, O(1) space
//func trap(height []int) int {
//	l, r := 0, len(height)-1
//	lMax := height[l]
//	rMax := height[r]
//	res := 0
//	for l < r {
//		if lMax < rMax {
//			l += 1
//			lMax = customMax(lMax, height[l])
//			res += lMax - height[l]
//		} else {
//			r -= 1
//			rMax = customMax(rMax, height[r])
//			res += rMax - height[r]
//		}
//	}
//
//	return res
//}
//
//func customMax(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// O(n) time O(n) space
func trap(height []int) int {
	l := len(height)
	lMaxArray := make([]int, l)
	rMaxArray := make([]int, l)
	resArray := make([]int, l)

	lMax := 0
	for i := 1; i < l-1; i++ {
		if height[i-1] > lMax {
			lMax = height[i-1]
		}
		lMaxArray[i] = lMax
	}

	rMax := 0
	for i := l - 2; i >= 1; i-- {
		if height[i+1] > rMax {
			rMax = height[i+1]
		}
		rMaxArray[i] = rMax
	}

	for i := 1; i < l-1; i++ {
		if lMaxArray[i] < rMaxArray[i] {
			resArray[i] = lMaxArray[i] - height[i]
		} else {
			resArray[i] = rMaxArray[i] - height[i]
		}
	}

	res := 0
	for i := 1; i < l-1; i++ {
		if resArray[i] > 0 {
			res += resArray[i]
		}
	}

	return res
}

func main() {
	testCase := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(testCase))
}
