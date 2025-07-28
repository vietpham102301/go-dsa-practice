package main

import "fmt"

func insertBits(N, M, i, j int) int {
	allOnes := ^0
	left := allOnes << (j + 1)
	right := (1 << i) - 1

	mask := left | right
	clearN := N & mask

	shiftedM := M << i

	res := clearN | shiftedM

	return res
}

func main() {
	N := 1024
	M := 19

	i := 2
	j := 6

	res := insertBits(N, M, i, j)
	fmt.Printf("Binary: %b\n", res)

}
