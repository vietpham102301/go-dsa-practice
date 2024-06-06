package main

import "fmt"

func permutation(s string) {
	permutationRecursive(s, "")
}

func permutationRecursive(s string, prefix string) {
	if len(s) == 0 {
		fmt.Println(prefix)
	} else {
		for i := 0; i < len(s); i++ {
			permutationRecursive(s[:i]+s[i+1:], prefix+string(s[i]))
		}
	}
}

func main() {
	permutation("abc")
}
