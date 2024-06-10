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

func checkPalindromePermutation2(s string) bool {
	var normalizeString strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			normalizeString.WriteByte(s[i])
		}
	}

	processedString := strings.ToLower(normalizeString.String())
	bitMap := createBitmap(processedString)

	return bitMap == 0 || checkExactlyOneBitSet(bitMap)
}

func createBitmap(s string) int {
	var bitMap int
	for _, ele := range s {
		if ele >= 'a' && ele <= 'z' {
			index := int(ele - 'a') // don't forget assuming a-z and get the index by minus 'a' char
			mask := 1 << index
			if bitMap&mask == 0 {
				bitMap |= mask
			} else {
				bitMap &= ^mask
			}
		}
	}
	return bitMap
}

func checkExactlyOneBitSet(bitMap int) bool {
	bitMapMinus1 := bitMap - 1

	return bitMap&bitMapMinus1 == 0
}

func main() {
	fmt.Println(checkPalindromePermutation2("Tact Coa"))
}
