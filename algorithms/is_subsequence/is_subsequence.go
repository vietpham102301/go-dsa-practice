package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	res := false
	sIndex := 0
	for i, _ := range t {
		if sIndex >= len(s) {
			break
		}
		if s[sIndex] == t[i] {
			sIndex++
			res = true
		} else {
			res = false
		}
	}

	return res && sIndex == len(s)
}

func main() {
	s := "b"
	t := "abc"

	res := isSubsequence(s, t)
	fmt.Println(res)
}
