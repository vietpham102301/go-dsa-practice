package main

import "strings"

func reverseWords(s string) string {
	trimmedStr := strings.TrimSpace(s)
	splitedStr := strings.Split(trimmedStr, " ")

	start := 0
	end := len(splitedStr) - 1

	for start < end {
		splitedStr[start], splitedStr[end] = splitedStr[end], splitedStr[start]

		start++
		end--
	}

	res := ""

	for i := 0; i < len(splitedStr); i++ {
		if splitedStr[i] != "" {
			res += splitedStr[i] + " "
		}
	}
	res = strings.TrimSpace(res)
	return res
}

func main() {
	s := "a good   example"
	reverseWords(s)
}
