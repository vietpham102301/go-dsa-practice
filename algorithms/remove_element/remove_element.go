package main

func removeElement(nums []int, val int) int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2} // Input array
	val := 2                              // Value to remove

	k := removeElement(nums, val) // Calls your implementation

	println(k)
}
