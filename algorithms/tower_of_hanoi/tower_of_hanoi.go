package main

import "fmt"

func towerOfHanoi(stepCount *int, n, source, target, auxiliary int) {
	*stepCount++
	if n == 1 {
		fmt.Printf("Move disk %d from %d to %d \n", n, source, target)
		return
	}

	towerOfHanoi(stepCount, n-1, source, auxiliary, target)
	fmt.Printf("Move disk %d from %d to %d \n", n, source, target)
	towerOfHanoi(stepCount, n-1, auxiliary, target, source)
}

func main() {
	stepCount := new(int)
	*stepCount = 0
	towerOfHanoi(stepCount, 5, 1, 3, 2)
	fmt.Printf("steps count: %d\n", *stepCount)
}
