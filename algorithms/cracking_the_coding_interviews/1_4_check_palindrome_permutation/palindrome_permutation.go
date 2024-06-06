package main

import (
	"fmt"
	"strings"
)

func checkPalindromePermutation(s string) bool {
	var normalizeString strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			normalizeString.WriteByte(s[i])
		}
	}

	processedString := strings.ToLower(normalizeString.String())
	hashTable := make(map[string]int)

	for i := 0; i < len(processedString); i++ {
		hashTable[string(processedString[i])]++
	}

	oddCount := 0
	for _, val := range hashTable {
		if val%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= 1
}

func main() {
	fmt.Println(checkPalindromePermutation("Tact Coa"))
}
