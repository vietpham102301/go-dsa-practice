package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func isBST(root *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	if root == nil {
		return true
	}

	if minNode != nil && root.data <= minNode.data || maxNode != nil && root.data >= maxNode.data {
		return false
	}

	return isBST(root.left, minNode, root) && isBST(root.right, root, maxNode)
}

func (node *TreeNode) Insert(value int) {
	if value < node.data {
		if node.left == nil {
			node.left = &TreeNode{data: value}
		} else {
			node.left.Insert(value)
		}
	} else {
		if node.right == nil {
			node.right = &TreeNode{data: value}
		} else {
			node.right.Insert(value)
		}
	}
}

func inOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	inOrderTraversal(node.left)
	fmt.Println(node.data)
	inOrderTraversal(node.right)
}

func main() {

	root := &TreeNode{data: 10}

	// values := []int{5, 15, 3, 7, 12, 18}
	// for _, value := range values {
	// 	root.Insert(value)
	// }

	root.left = &TreeNode{data: 7}
	root.right = &TreeNode{data: 12}
	root.left.left = &TreeNode{data: 5}
	root.left.right = &TreeNode{data: 11}

	inOrderTraversal(root)

	fmt.Println(isBST(root, nil, nil))
}
