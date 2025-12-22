package main

import "fmt"

var (
	xi = []int{2, 1, 2, 1, -2, -1, -2, -1}
	yi = []int{1, 2, -1, -2, -1, -2, 1, 2}
)

func knightsTour(x int, y int, step int, board [][]int) bool {
	boardSize := len(board)
	board[x][y] = step

	if step == boardSize*boardSize {
		return true
	}

	for i := 0; i < 8; i++ {
		xNext := x + xi[i]
		yNext := y + yi[i]

		if isValid(xNext, yNext, board) {
			if knightsTour(xNext, yNext, step+1, board) {
				return true
			}
		}
	}

	board[x][y] = 0
	return false
}

func isValid(x, y int, board [][]int) bool {
	size := len(board)
	return x >= 0 && y >= 0 && x < size && y < size && board[x][y] == 0
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
			fmt.Printf("| %2d ", board[i][j])
		}
		fmt.Println("|")
	}
}

func main() {
	size := 8
	board := initChessBoard(size)

	// Bắt đầu tại vị trí (0, 0) hoặc (2, 2) với bước đi đầu tiên là 1
	startX, startY := 2, 2

	if knightsTour(startX, startY, 1, board) {
		fmt.Println("Đã tìm thấy đường đi:")
		printChessBoard(board)
	} else {
		fmt.Println("Không tìm thấy giải pháp.")
	}
}
