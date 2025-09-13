package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	maxLength := 0
	l := 0
	cachedMap := make(map[uint8]bool)
	for r := 0; r < len(s); r++ {
		for cachedMap[s[r]] {
			delete(cachedMap, s[l])
			l++
		}
		cachedMap[s[r]] = true
		if r-l+1 > maxLength {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
