package main

import "fmt"

func modifySlice(smt []int) {
	smt[2] = 123
}

func main() {
	test := []int{1, 2, 3, 4}
	modifySlice(test)
	for _, val := range test {
		fmt.Printf("%d, ", val)
	}
}
