package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	trimedStr := strings.TrimSpace(s)
	splitedStr := strings.Split(trimedStr, " ")

	lastWord := splitedStr[len(splitedStr)-1]

	return len(lastWord)
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
