package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	rowIndex := 0
	direction := 1
	for _, char := range s {
		rows[rowIndex] += string(char)
		if rowIndex == 0 {
			direction = 1
		} else if rowIndex == numRows-1 {
			direction = -1
		}
		rowIndex += direction
	}

	ret := ""
	for _, row := range rows {
		ret += row
	}

	return ret
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	fmt.Println(convert(s, numRows))
}
