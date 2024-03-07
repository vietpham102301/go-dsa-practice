package main

import "fmt"

func productExceptSelf(nums []int) []int {
	l := len(nums) - 1
	prefixProduct := make([]int, l+1)
	suffixProduct := make([]int, l+1)
	//init
	for i := 0; i < l+1; i++ {
		prefixProduct[i] = 1
		suffixProduct[i] = 1
	}

	for i := 1; i < l+1; i++ {
		prefixProduct[i] = prefixProduct[i-1] * nums[i-1]
	}

	for i := l - 1; i >= 0; i-- {
		suffixProduct[i] = suffixProduct[i+1] * nums[i+1]
	}

	ans := make([]int, l+1)

	for i := 0; i < l+1; i++ {
		ans[i] = prefixProduct[i] * suffixProduct[i]
	}
	return ans
}

func main() {
	testCase := []int{1, 2, 3, 4}

	res := productExceptSelf(testCase)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
