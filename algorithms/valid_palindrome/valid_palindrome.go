package main

import "fmt"

func isPalindromeValid(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for !isAlphanumerical(s[left]) {
			left++
		}

		for !isAlphanumerical(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumerical(char uint8) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func main() {
	testCases := []string{"abcde", "racecar", "123321", "s ! t ! # $ t   s ", "hello, world", ""}

	for _, test := range testCases {
		fmt.Println(isPalindromeValid(test))
	}
}
