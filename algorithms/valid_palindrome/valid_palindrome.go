package main

import "fmt"

func isPalindrome(s string) bool {
	s = removeNonAlphanumericChars(s)
	return checkPalindrome(s)
}

func removeNonAlphanumericChars(s string) string {
	ret := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			ret += string(char)
		}
		if char >= 'A' && char <= 'Z' { //convert to lower case letter
			ret += string(char + 32)
		}
	}
	fmt.Println(ret)
	return ret
}

func checkPalindrome(s string) bool {
	forwardPointer := 0
	backwardPointer := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[forwardPointer] != s[backwardPointer] {
			return false
		}
		forwardPointer++
		backwardPointer--
	}

	return true
}

func isPalindromeTwoPointer(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !isAlphanumeric(runes[left]) {
			left++
		}
		for left < right && !isAlphanumeric(runes[right]) {
			right--
		}
		leftChar := toLower(runes[left])
		rightChar := toLower(runes[right])
		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func main() {
	s := "race a car"
	valid := isPalindromeTwoPointer(s)
	fmt.Println(valid)
}
