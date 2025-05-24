package main

import "fmt"

func zeroStriping(arr [][]int) [][]int {
	m, n := len(arr), len(arr[0])
	if m == 0 || n == 0 {
		return arr
	}

	rowSet := make(map[int]bool, m)
	colSet := make(map[int]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				rowSet[i] = true
				colSet[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowSet[i] || colSet[j] {
				arr[i][j] = 0
			}
		}
	}

	return arr
}

func zeroStripingNoExtraSpace(arr [][]int) [][]int {
	m, n := len(arr), len(arr[0])
	if m == 0 || n == 0 {
		return arr
	}

	firstRowZero := false
	firstColZero := false

	for i := 0; i < n; i++ {
		if arr[0][i] == 0 {
			firstRowZero = true
		}
	}

	for i := 0; i < m; i++ {
		if arr[i][0] == 0 {
			firstColZero = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				arr[0][j] = 0
				arr[i][0] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[0][j] == 0 || arr[i][0] == 0 {
				arr[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for i := 0; i < n; i++ {
			arr[0][i] = 0
		}
	}

	if firstColZero {
		for i := 0; i < m; i++ {
			arr[i][0] = 0
		}
	}

	return arr
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

func copyMatrix(matrix [][]int) [][]int {
	result := make([][]int, len(matrix))
	for i := range matrix {
		result[i] = make([]int, len(matrix[i]))
		copy(result[i], matrix[i])
	}
	return result
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println("Original Matrix:")
	printMatrix(input)

	fmt.Println("Using zeroStriping:")
	result1 := zeroStriping(copyMatrix(input))
	printMatrix(result1)

	fmt.Println("Using zeroStripingNoExtraSpace:")
	result2 := zeroStripingNoExtraSpace(copyMatrix(input))
	printMatrix(result2)

}
