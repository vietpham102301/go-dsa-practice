package main

import "fmt"

type HashSet struct {
	arr map[int]bool
}

func (h *HashSet) add(element int) bool {
	_, ok := h.arr[element]
	if !ok {
		h.arr[element] = true
		return true
	}
	return false
}

func InitHashSet(cap int) *HashSet {
	return &HashSet{
		arr: make(map[int]bool, cap),
	}
}

func InitSubgridSet(cap int) [][]*HashSet {
	subgridSet := make([][]*HashSet, cap)
	for i := 0; i < cap; i++ {
		subgridSet[i] = make([]*HashSet, cap)
		for j := 0; j < cap; j++ {
			subgridSet[i][j] = InitHashSet(3)
		}
	}
	return subgridSet
}

func InitArrSet(cap int) []*HashSet {
	arrSet := make([]*HashSet, 9)
	for i := 0; i < cap; i++ {
		arrSet[i] = InitHashSet(9)
	}

	return arrSet
}

func verifySudokuBoard(board [9][9]int) bool {
	rowSet := InitArrSet(9)
	colSet := InitArrSet(9)
	subgridSet := InitSubgridSet(3)

	for i, _ := range board {
		for j, _ := range board[i] {
			num := board[i][j]
			if num == 0 {
				continue // we didn't take 0 into account
			}

			ok := rowSet[i].add(num) || colSet[j].add(num)
			if !ok {
				return false
			}

			ok2 := subgridSet[i/3][j/3].add(num)
			if !ok2 {
				return false
			}
		}
	}

	return true
}

func main() {
	testCase := [9][9]int{
		{3, 0, 6, 0, 5, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{1, 0, 2, 5, 0, 0, 3, 2, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{0, 3, 0, 0, 0, 8, 2, 5, 0},
		{0, 1, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 0, 0, 0},
	}

	res := verifySudokuBoard(testCase)
	fmt.Println(res)
}
