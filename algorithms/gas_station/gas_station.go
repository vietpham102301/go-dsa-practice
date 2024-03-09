package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; i++ {
		tank := 0
		j := i
		secondTime := false
		itCan := true
		for j != i || !secondTime {
			tank = tank + gas[j] - cost[j]
			j++
			if j > l-1 {
				j = j % l
			}
			if j == i && tank >= 0 {
				secondTime = true
				continue
			}
			if tank <= 0 {
				itCan = false
				break
			}
		}
		if !itCan {
			continue
		}
		return i
	}
	return -1
}

func main() {
	gas := []int{4, 5, 3, 1, 4}
	cost := []int{5, 4, 3, 4, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
