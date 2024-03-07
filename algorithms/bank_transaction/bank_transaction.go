package main

import (
	"fmt"
	"math"
)

// The question is how many transactions could be made
func maximizeTransaction(transactions []int) int {
	var maxTrans []int
	maxTrans[0] = 1

	for i := 1; i < len(transactions); i++ {
		sum := transactions[0]
		maxTrans[i] = math.MinInt
		for j := 0; j <= i; j++ {
			sum += transactions[j]
			if sum >= 0 {
				maxTrans[i] = customMax(maxTrans[i], maxTrans[j]+1)
			}
		}
	}

	return maxTrans[len(transactions)-1]
}

func customMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	transactions := []int{3, 2, -5, -6, -1, 4}
	result := maximizeTransaction(transactions)
	fmt.Println(result)
}
