package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := 0
	minLength := len(nums) + 1
	for _, num := range nums {
		sum += num
		for sum >= target {
			if minLength > r-l+1 {
				minLength = r - l + 1
			}
			sum -= nums[l]
			l++
		}
		r++
	}
	if minLength == len(nums)+1 {
		return 0
	}

	return minLength
}

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7
	res := minSubArrayLen(target, arr)
	fmt.Println(res)
}
