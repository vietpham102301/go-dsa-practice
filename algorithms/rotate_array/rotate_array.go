package main

import "fmt"

// my stupid way takes O(k*n) but k is always a constant and can be
// modulo so is can be simplified to O(n) but using extra space to store secondPart
func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) <= 1 {
		return
	}
	secondPart := make([]int, k)
	copy(secondPart, nums[len(nums)-k:])
	shiftByK(nums, k)
	for i := 0; i < k; i++ {
		nums[i] = secondPart[i]
	}
}

func shiftByK(nums []int, k int) {
	for i := 0; i < k; i++ {
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
	}
}

// solution way takes O(n)
//func reverse(nums []int, start, end int) {
//	for start < end {
//		nums[start], nums[end] = nums[end], nums[start]
//		start++
//		end--
//	}
//}
//
//func rotate(nums []int, k int) {
//	n := len(nums)
//	k %= n
//
//	reverse(nums, 0, n-1)
//	reverse(nums, 0, k-1)
//	reverse(nums, k, n-1)
//}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
