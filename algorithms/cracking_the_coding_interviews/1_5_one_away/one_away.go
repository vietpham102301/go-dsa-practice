package main

import (
	"fmt"
	"math"
)

func checkOneEdit(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	// if insert or remove more than once
	removeOrInsertTime := math.Abs(float64(l1 - l2))
	if removeOrInsertTime >= 2.0 {
		return false
	}

	if removeOrInsertTime == 1 {
		hashMap := make(map[uint8]bool)

		if l1 < l2 {
			for i := 0; i < l2; i++ {
				hashMap[s2[i]] = true
			}

			for i := 0; i < l1; i++ {
				val := hashMap[s1[i]]
				if !val {
					return false
				}
			}
		} else {
			for i := 0; i < l1; i++ {
				hashMap[s1[i]] = true
			}

			for i := 0; i < l2; i++ {
				val := hashMap[s2[i]]
				if !val {
					return false
				}
			}
		}
	} else if removeOrInsertTime == 0 {
		count := 0
		for i := 0; i < l1; i++ {
			if s1[i] != s2[i] {
				count++
			}
		}
		if count >= 2 {
			return false
		}
	}

	return true
}

func checkOneEdit2(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)

	// If the length difference is more than 1, return false
	if math.Abs(float64(l1-l2)) > 1 {
		return false
	}

	// Helper function to check for one replacement
	checkReplace := func(s1, s2 string) bool {
		foundDifference := false
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if foundDifference {
					return false
				}
				foundDifference = true
			}
		}
		return true
	}

	// Helper function to check for one insertion/removal
	checkInsertRemove := func(s1, s2 string) bool {
		index1, index2 := 0, 0
		for index1 < len(s1) && index2 < len(s2) {
			if s1[index1] != s2[index2] {
				if index1 != index2 {
					return false
				}
				index2++
			} else {
				index1++
				index2++
			}
		}
		return true
	}

	if l1 == l2 {
		return checkReplace(s1, s2)
	} else if l1+1 == l2 {
		return checkInsertRemove(s1, s2)
	} else if l1 == l2+1 {
		return checkInsertRemove(s2, s1)
	}

	return false
}

func main() {
	fmt.Println(checkOneEdit("pale", "bake"))

}
