package main

import "fmt"

func towerOfHanoi(count *int, diskNum int, source, target, sub int) {
	*count++
	if diskNum == 1 {
		printMove(diskNum, source, target)
		return
	}

	towerOfHanoi(count, diskNum-1, source, sub, target)
	printMove(diskNum, source, target)

	towerOfHanoi(count, diskNum-1, sub, target, source)
}

func printMove(disk, from, to int) {
	fmt.Printf("move disk %d from %d to %d\n", disk, from, to)
}

func main() {
	count := 0
	towerOfHanoi(&count, 3, 1, 3, 2)
	fmt.Printf("total moves: %d \n", count)
}
