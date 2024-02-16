package main

import "fmt"

func maxProfit(prices []int) int {
	accumulateProfit := 0
	i := 0
	j := 1
	for j < len(prices) {
		profit := prices[j] - prices[i]
		if profit > 0 {
			accumulateProfit += profit
		}
		i++
		j++
	}

	return accumulateProfit
}

//Intuition solutions
//func maxProfit(prices []int) int {
//	accumulateProfit := 0
//	i := 0
//	j := 1
//
//	for j < len(prices) {
//		profit := prices[j] - prices[i]
//		if profit > 0 {
//			accumulateProfit += profit
//			i++
//			j++
//		}else {
//			i = j
//			j++
//		}
//	}
//
//	return accumulateProfit
//
//}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	accProfit := maxProfit(prices)

	fmt.Println(accProfit)
}
