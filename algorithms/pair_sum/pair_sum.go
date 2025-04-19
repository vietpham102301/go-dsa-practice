package main

import "fmt"

func pairSum(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {
	testcase := []int{1, 2, 3, 4, 5}
	target := 6

	res := pairSum(testcase, target)

	fmt.Println(res)
}
