package main

import (
	"fmt"
	"math"
)

func sieveOfEratosthenes(max int) []bool {
	prime := 2
	flags := make([]bool, max)
	initFlag(&flags)
	for float64(prime) <= math.Sqrt(float64(max)) {
		crossOff(prime, &flags)
		prime = getNextPrime(prime, flags)
	}

	return flags
}

func initFlag(flags *[]bool) {
	for i := 2; i < len(*flags); i++ {
		(*flags)[i] = true
	}
}

func crossOff(prime int, flags *[]bool) {
	for i := prime * prime; i < len(*flags); i += prime {
		(*flags)[i] = false
	}
}

func getNextPrime(prime int, flags []bool) int {
	next := prime + 1
	for next < len(flags) && !flags[prime] {
		next++
	}

	return next
}

func main() {
	res := sieveOfEratosthenes(100)

	for i, f := range res {
		if f {
			fmt.Printf("%d, ", i)
		}
	}
}
