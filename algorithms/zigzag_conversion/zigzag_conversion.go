package main

import "fmt"

func convert(s string, numRows int) string {
	zigZagStr := [][]string{}
	res := ""

	tempStr := []string{}

	for i := 0; i < len(s); i++ {
		tempStr = append(tempStr, string(s[i]))
		if len(tempStr) == numRows {
			zigZagStr = append(zigZagStr, tempStr)
			tempStr = []string{}
			i--
		}
	}
	zigZagStr = append(zigZagStr, tempStr)

	for i := 0; i < numRows; i++ {
		for j, strs := range zigZagStr {
			if (j%2) == 0 && i%2 == 0 {
				if i <= len(strs)-1 {
					res += strs[i]
				}
			}
			if i%2 != 0 {
				if i <= len(strs)-1 {
					res += strs[i]
				}
			}
		}
	}

	return res
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	fmt.Println(convert(s, numRows))
}
