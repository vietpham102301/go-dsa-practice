package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringCompression(s string) string {
	result := strings.Builder{}
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		if i+1 >= len(s) || s[i+1] != s[i] {
			result.WriteRune(rune(s[i]))
			result.WriteString(strconv.Itoa(count))
			count = 0
		}
	}

	if len(result.String()) >= len(s) {
		return s
	} else {
		return result.String()
	}
}

func main() {
	s := "aabcccccaaa"

	fmt.Println(stringCompression(s))
}
