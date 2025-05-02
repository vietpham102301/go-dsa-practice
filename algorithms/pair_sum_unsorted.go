package main

import "fmt"

func pairSumUnsorted(arr []int, target int) []int {
	numMap := make(map[int]int)
	for i, ele := range arr {
		j, ok := numMap[target-ele]
		if ok {
			return []int{i, j}
		}

		numMap[ele] = i
	}

	return []int{}
}

func main() {
	testCases := [][]int{
		{},
		{1, 2, 3},
		{4, -1, 2},
		{3, 4, 5},
		{-2, 0, 5, 3},
	}

	for _, testCase := range testCases {
		res := pairSumUnsorted(testCase, 3)
		fmt.Println(res)
	}

}
