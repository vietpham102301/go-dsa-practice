package main

import "fmt"

func rotateMatrix(image [][]int) [][]int {
	n := len(image[0])
	var result [][]int
	for i := 0; i < n; i++ {
		var temp []int
		for j := n - 1; j >= 0; j-- {
			temp = append(temp, image[j][i])
		}
		result = append(result, temp)
	}

	return result
}

func rotateMatrix2(image [][]int) [][]int {
	n := len(image)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = image[n-1-j][i]
		}
	}

	return result
}

func rotateMatrixInPlace(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) != len(matrix) {
		return nil
	}
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ { // i increase from first
			idx := i - first //index that increase from 0

			//save the top
			top := matrix[first][i]

			//rotate
			matrix[first][i] = matrix[last-idx][first]
			matrix[last-idx][first] = matrix[last][last-idx]
			matrix[last][last-idx] = matrix[i][last]
			matrix[i][last] = top //use saved top
		}
	}

	return matrix
}

func main() {
	image := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(rotateMatrixInPlace(image))
}
