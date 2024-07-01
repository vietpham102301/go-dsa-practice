package main

import "fmt"

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode // case we have link to parent
	data   int
}

func findFirstCommonAncestor(node1, node2 *TreeNode) *TreeNode {
	depth1 := getDepth(node1)
	depth2 := getDepth(node2)

	delta := depth1 - depth2

	var first, second *TreeNode
	if delta > 0 {
		first = node1
		second = node2
	} else {
		first = node2
		second = node1
	}

	first = goUpBy(first, delta)

	for first != nil && second != nil && first != second {
		first = first.parent
		second = second.parent
	}

	if first == nil || second == nil {
		return nil
	} else {
		return first
	}
}

func getDepth(node *TreeNode) int {
	curr := node
	depth := 0
	for curr != nil {
		depth++
		curr = curr.parent
	}

	return depth
}

func goUpBy(node *TreeNode, delta int) *TreeNode {
	for delta != 0 && node != nil {
		node = node.parent
		delta--
	}

	return node
}

func findFirstCommonAncestor2(root, node1, node2 *TreeNode) *TreeNode {
	if !isCover(root, node1) || !isCover(root, node2) {
		return nil
	} else if isCover(node1, node2) {
		return node1
	} else if isCover(node2, node1) {
		return node2
	}

	// go upwards
	parent := node1.parent
	sibling := getSibling(node1)

	for !isCover(sibling, node2) {
		sibling = getSibling(parent)
		parent = parent.parent
	}

	return parent
}

func isCover(root, node *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}

	return isCover(root.left, node) || isCover(root.right, node)
}

func getSibling(node *TreeNode) *TreeNode {
	if node == nil || node.parent == nil {
		return nil
	}

	parent := node.parent

	if parent.left == node {
		return parent.right
	} else {
		return parent.left
	}
}

func main() {
	root := &TreeNode{data: 20}
	root.left = &TreeNode{data: 10}
	root.left.parent = root
	root.right = &TreeNode{data: 30}
	root.right.parent = root

	root.left.left = &TreeNode{data: 5}
	root.left.left.parent = root.left
	root.left.right = &TreeNode{data: 15}
	root.left.right.parent = root.left

	root.left.right.right = &TreeNode{data: 17}
	root.left.right.right.parent = root.left.right

	root.left.left.left = &TreeNode{data: 3}
	root.left.left.left.parent = root.left.left.left
	root.left.left.right = &TreeNode{data: 7}
	root.left.left.right.parent = root.left.left

	res := findFirstCommonAncestor2(root, root.left.right.right, root.left.left.right)

	fmt.Println(res.data)
}
