package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if s == "" {
		return []int{}
	}
	permutedStrings := permuteStrings(words)

	concatedStrings := []string{}
	for _, p := range permutedStrings {
		temp := ""
		for _, s := range p {
			temp += s
		}
		concatedStrings = append(concatedStrings, temp)
	}

	l := len(concatedStrings[0])

	cachedMap := make(map[string]int)

	for i, _ := range s {
		if i+l > len(s) {
			break
		}
		temp := s[i : i+l]
		cachedMap[temp] = i
	}

	res := []int{}
	duplicateMap := make(map[int]bool)

	for _, s := range concatedStrings {
		index, ok := cachedMap[s]
		if ok && !duplicateMap[index] {
			res = append(res, index)
			duplicateMap[index] = true
		}
	}

	return res
}

func permuteStrings(arr []string) [][]string {
	var result [][]string
	var backtrack func(int)

	backtrack = func(l int) {
		if l == len(arr) {
			perm := make([]string, len(arr))
			copy(perm, arr)
			result = append(result, perm)
			return
		}
		for i := l; i < len(arr); i++ {
			arr[l], arr[i] = arr[i], arr[l]
			backtrack(l + 1)
			arr[l], arr[i] = arr[i], arr[l] // backtrack
		}
	}

	backtrack(0)
	return result
}

func main() {
	s := "foobarfoobar"
	words := []string{"foo", "bar"}

	res := findSubstring(s, words)
	fmt.Println(res)
}
