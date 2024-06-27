package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != math.MinInt
}

func checkHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}

	leftHeight := checkHeight(root.left)
	if leftHeight == math.MinInt {
		return math.MinInt
	}
	rightHeight := checkHeight(root.right)
	if rightHeight == math.MinInt {
		return math.MinInt
	}

	heightDiff := rightHeight - leftHeight
	if math.Abs(float64(heightDiff)) > 1 {
		return math.MinInt
	} else {
		return max(leftHeight, rightHeight) + 1
	}
}

func main() {
	root := &TreeNode{data: 1}
	root.left = &TreeNode{data: 2}
	root.right = &TreeNode{data: 3}
	root.left.left = &TreeNode{data: 4}
	root.left.right = &TreeNode{data: 5}
	// root.left.left.left = &TreeNode{data: 6}

	fmt.Println(isBalanced(root))
}
