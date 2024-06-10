package main

import "fmt"

type Point struct {
	x int
	y int
}

func mutateMatrix(matrix [][]int, m int, n int) [][]int {
	var zeroPoints []Point
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroPoints = append(zeroPoints, Point{i, j})
			}
		}
	}

	for _, point := range zeroPoints {
		zerolize(matrix, point.x, point.y, m, n)
	}

	return matrix
}

func zerolize(matrix [][]int, x int, y int, m int, n int) {
	for i := 0; i < n; i++ {
		matrix[x][i] = 0
	}
	for i := 0; i < m; i++ {
		matrix[i][y] = 0
	}
}

func main() {

	image := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 0, 12},
	}

	fmt.Println(mutateMatrix(image, 3, 4))
}
