package main

import "fmt"

func isSubstring(word1, word2 string) int {
	l1, l2 := len(word1), len(word2)
	for i := 0; i < l2; i++ {
		if word2[i] == word1[0] {
			index := i
			for j := 0; j < l1; j++ {
				if index >= l2 {
					return -1
				}
				if word1[j] == word2[index] {
					index++
				} else {
					break
				}
			}
			if index-i == len(word1) {
				return i
			}
		}
	}

	return -1
}

/*
	We can see that, we always can find a xy such xy = s1 and yx = s2:
	s1 := xy = waterbottle
	x = wat
	y = erbottle
	s2 := yx = erbottlewat
	yx always be a substring of xyxy
	=> s2(yx) is always be a substring of s1s1(xyxy)
*/

func checkRotation(s1, s2 string) bool {
	s1s1 := s1 + s1

	startIndex := isSubstring(s2, s1s1)

	if startIndex != -1 {
		return true
	}

	return false
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	startIndex := checkRotation(s2, s1)

	fmt.Println(startIndex)
}
