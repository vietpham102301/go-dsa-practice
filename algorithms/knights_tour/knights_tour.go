package main

import "fmt"

func knightsTour(x int, y int, board [][]int, step int, q *bool) {
	xi := []int{2, 1, 2, 1, -2, -1, -2, -1}
	yi := []int{1, 2, -1, -2, -1, -2, 1, 2}

	boardSize := len(board)

	if step == boardSize*boardSize {
		*q = true
		return
	}

	for i := 0; (i < len(xi)) && (*q == false); i++ {
		xNext := x + xi[i]
		yNext := y + yi[i]
		if xNext >= 0 && yNext >= 0 && xNext < boardSize && yNext < boardSize && board[xNext][yNext] == 0 {
			board[xNext][yNext] = step
			knightsTour(xNext, yNext, board, step+1, q)
			if !(*q) {
				board[xNext][yNext] = 0
			}
		}
	}
}

func initChessBoard(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
	}

	return res
}

func printChessBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("| - %d - |", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := initChessBoard(5)

	q := new(bool)

	board[2][2] = 1
	knightsTour(2, 2, board, 2, q)

	printChessBoard(board)
}
