package main

import "fmt"

func fibRecursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

func printFib(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(fibonacciMemoization(i))
	}
}

var memo [1000]int

func fibonacciMemoization(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = fibonacciMemoization(n-1) + fibonacciMemoization(n-2)

	return memo[n]
}

func main() {
	printFib(10)
}
