package main

import "fmt"

func isUnique(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func isUnique2(s string) bool {
	l := len(s)

	if l > 128 {
		return false
	}
	var visited [128]bool

	for i := 0; i < l; i++ {
		if visited[int(s[i])] {
			return false
		}

		visited[int(s[i])] = true
	}

	return true
}

func isUnique3(s string) bool {
	var checker int

	for i := 0; i < len(s); i++ {
		val := s[i] - 'a'
		if checker&(1<<val) > 0 {
			return false
		}

		checker |= 1 << val
	}

	return true
}

func main() {
	s := "abcdefgh"
	fmt.Println(isUnique3(s))
}
