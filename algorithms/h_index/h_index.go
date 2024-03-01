package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	h := -1
	for i := 1; i <= len(citations); i++ {
		count := 0
		for j := 0; j < i; j++ {
			if citations[j] >= i {
				count++
			}
		}
		if count > h {
			h = count
		}
	}
	return h
}

func main() {
	testCase := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(testCase))
}
