package main

func removeElementOptimize(nums []int, val int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

func removeElementIntuition(nums []int, val int) int {
	var indiesOfVal []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			indiesOfVal = append(indiesOfVal, i)
		}
	}
	var lastIndiesToReplace []int
	for i := len(nums) - 1; i >= 0; i-- {
		isContains := false
		for j := len(indiesOfVal) - 1; j >= 0; j-- {
			if i == indiesOfVal[j] {
				isContains = true
			}
		}
		if !isContains && len(lastIndiesToReplace) != len(indiesOfVal) {
			lastIndiesToReplace = append(lastIndiesToReplace, i)
		}
	}

	for i := 0; i < len(nums)-len(indiesOfVal); i++ {
		if nums[i] == val {
			for j := 0; j < len(lastIndiesToReplace); j++ {
				nums[i], nums[lastIndiesToReplace[j]] = nums[lastIndiesToReplace[j]], nums[i]
				lastIndiesToReplace = lastIndiesToReplace[j+1:]
				break
			}
		}
	}
	return len(nums) - len(indiesOfVal)
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2} // Input array
	val := 2                              // Value to remove

	k := removeElementOptimize(nums, val) // Calls your implementation

	println(k)
}
