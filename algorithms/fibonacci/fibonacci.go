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
	for i := 0; i <= n; i++ {
		fmt.Println(fibonacciIter(i))
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

// 0-index base fibonacci
func fibonacciIter(n int) int {
	first := 0
	second := 1
	if n == 0 {
		return first
	} else if n == 1 {
		return second
	}

	var fib int
	for i := 2; i <= n; i++ {
		fib = first + second
		first = second
		second = fib
	}

	return fib
}

func main() {
	printFib(10)
}
