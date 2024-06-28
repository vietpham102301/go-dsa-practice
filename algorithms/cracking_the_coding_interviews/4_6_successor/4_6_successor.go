package main

import "fmt"

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	data   int
}

func (n *TreeNode) findMin() *TreeNode {
	curr := n
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (n *TreeNode) findSuccessor() *TreeNode {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return n.right.findMin()
	}

	curr := n
	parent := n.parent

	for parent != nil && curr == parent.right {
		curr = parent
		parent = parent.parent
	}

	return parent
}

func main() {
	root := &TreeNode{data: 20}
	node10 := &TreeNode{data: 10, parent: root}
	node30 := &TreeNode{data: 30, parent: root}
	node5 := &TreeNode{data: 5, parent: node10}
	node15 := &TreeNode{data: 15, parent: node10}
	node25 := &TreeNode{data: 25, parent: node30}
	node35 := &TreeNode{data: 35, parent: node30}

	root.left = node10
	root.right = node30
	node10.left = node5
	node10.right = node15
	node30.left = node25
	node30.right = node35

	successor := node10.findSuccessor()

	fmt.Println(successor.data)

}
