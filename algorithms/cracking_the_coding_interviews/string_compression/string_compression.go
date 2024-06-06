package main

import (
	"fmt"
	"strconv"
)

func stringCompression(s string) string {
	result := ""
	for i := 0; i < len(s); {
		count := 0
		letter := s[i]
		for j := i; j < len(s); j++ {
			if letter == s[j] {
				count++
				i++
			} else {
				break
			}
		}
		result += string(letter) + strconv.Itoa(count)
	}

	if len(result) >= len(s) {
		return s
	} else {
		return result
	}
}

func main() {
	s := "aabcccccaaa"

	fmt.Println(stringCompression(s))
}
