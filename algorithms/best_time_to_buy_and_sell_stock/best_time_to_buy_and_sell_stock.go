package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	i := 0
	j := 1
	for j < len(prices) {
		profit := prices[j] - prices[i]
		if profit >= 0 {
			if profit > maxProfit {
				maxProfit = profit
			}
			j++
		} else {
			i = j
			j++
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	maxProfit := maxProfit(prices)

	fmt.Println(maxProfit)
}
