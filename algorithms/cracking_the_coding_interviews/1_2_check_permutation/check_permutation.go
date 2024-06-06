package main

import "fmt"

func checkPermutation(s1, s2 string) (bool, error) {
	if len(s1) != len(s2) {
		return false, nil
	}

	hashTable := make(map[rune]int)
	for _, r := range s1 {
		hashTable[r]++
	}
	for _, r := range s2 {
		hashTable[r]--
	}

	for _, count := range hashTable {
		if count != 0 {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	fmt.Println(checkPermutation("abbc", "cdba"))
}
