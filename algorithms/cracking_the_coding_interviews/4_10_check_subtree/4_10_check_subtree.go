package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func checkSubtree(t1 *TreeNode, t2 *TreeNode) bool {
	var s1 string
	var s2 string

	getStringFromTree(t1, &s1)
	getStringFromTree(t2, &s2)

	return strings.Contains(s1, s2)
}

func getStringFromTree(node *TreeNode, s *string) {
	if node == nil {
		*s += "X"
		return
	}

	*s += strconv.Itoa(node.data)
	getStringFromTree(node.left, s)
	getStringFromTree(node.right, s)
}

func checkSubtree2(t1 *TreeNode, t2 *TreeNode) bool {
	if t2 == nil {
		return true
	}

	return subTree(t1, t2)
}

func subTree(t1, t2 *TreeNode) bool {
	if t1 == nil {
		return false
	} else if t1.data == t2.data && matchTree(t1, t2) {
		return true
	}

	return subTree(t1.left, t2) || subTree(t1.right, t2)
}

func matchTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.data != t2.data {
		return false
	} else { // match data and recurse to base case
		return matchTree(t1.left, t2.left) && matchTree(t1.right, t2.right)
	}

}

func main() {
	root := &TreeNode{data: 32}
	root.left = &TreeNode{data: 12}
	root.right = &TreeNode{data: 37}
	root.left.left = &TreeNode{data: 9}
	root.left.right = &TreeNode{data: 15}
	root.right.left = &TreeNode{data: 33}
	root.right.right = &TreeNode{data: 40}

	root2 := &TreeNode{data: 37}
	root2.left = &TreeNode{data: 33}
	root2.right = &TreeNode{data: 40}

	fmt.Println(checkSubtree2(root, root2))
}
