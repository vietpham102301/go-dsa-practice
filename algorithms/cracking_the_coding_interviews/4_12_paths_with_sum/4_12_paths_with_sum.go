package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func countPathsWithSum(root *Node, targetSum int) int {
	return countPaths(root, make(map[int]int), 0, targetSum)
}

func countPaths(root *Node, pathCount map[int]int, runningSum int, targetSum int) int {
	if root == nil {
		return 0
	}
	runningSum += root.value

	s := runningSum - targetSum
	totalPaths := pathCount[s]
	if runningSum == targetSum {
		totalPaths++
	}
	incrementPathCount(pathCount, runningSum, 1)
	totalPaths += countPaths(root.left, pathCount, runningSum, targetSum)
	totalPaths += countPaths(root.right, pathCount, runningSum, targetSum)
	incrementPathCount(pathCount, runningSum, -1)

	return totalPaths
}

func incrementPathCount(pathCount map[int]int, runningSum, delta int) {
	newCount := pathCount[runningSum] + delta
	if newCount == 0 {
		delete(pathCount, runningSum)
	} else {
		pathCount[runningSum] = newCount
	}
}

func main() {
	/*
	         10
	        /  \
	       5   -3
	      / \    \
	     3   2    11
	    / \   \
	   3  -2   1
	*/

	root := &Node{value: 10}
	root.left = &Node{value: 5}
	root.right = &Node{value: -3}

	root.left.left = &Node{value: 3}
	root.left.right = &Node{value: 2}
	root.right.right = &Node{value: 11}

	root.left.left.left = &Node{value: 3}
	root.left.left.right = &Node{value: -2}
	root.left.right.right = &Node{value: 1}

	targetSum := 8
	fmt.Printf("Number of paths with sum %d: %d\n", targetSum, countPathsWithSum(root, targetSum))
}
