package main

import "fmt"

func testPassingSlice(s *[]int) {
	*s = append(*s, 1023)
}

func main() {
	originSlice := []int{1, 2}

	testPassingSlice(&originSlice)

	fmt.Println(originSlice)
}
