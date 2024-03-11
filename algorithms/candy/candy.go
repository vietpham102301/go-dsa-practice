package main

import "fmt"

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	acc := 0
	for i := 0; i < len(ratings); i++ {
		acc += candies[i]
	}
	return acc
}

func main() {
	test := []int{1, 6, 10, 8, 7, 3, 2}
	fmt.Println(candy(test))
}
