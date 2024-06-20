package main

import "fmt"

func createAdjacencyMatrix(n int, edges [][2]int) [][]int {
	res := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		res[i] = make([]int, n+1)
	}

	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		res[x][y] = 1
		res[y][x] = 1
	}

	return res
}

func printMatrix(matrix [][]int) {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix); j++ {
			fmt.Printf("%d, ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	n := 5
	edges := [][2]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{3, 5},
	}

	adjMatrix := createAdjacencyMatrix(n, edges)

	printMatrix(adjMatrix)
}
