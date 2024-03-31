package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	shortestStr := ""
	min := 100000
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min {
			shortestStr = strs[i]
			min = len(strs[i])
		}
	}

	longestPrefix := ""

	for i, c := range shortestStr {
		isAllHaveChar := true
		for _, str := range strs {
			if string(str[i]) != string(c) {
				isAllHaveChar = false
				break
			}
		}

		if !isAllHaveChar {
			break
		}

		if isAllHaveChar {
			longestPrefix += string(c)
		}
	}

	return longestPrefix
}

func main() {
	test := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(test))
}
