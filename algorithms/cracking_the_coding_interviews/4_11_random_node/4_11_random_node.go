package main

import (
	"fmt"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func (t *TreeNode) Insert(root *TreeNode, data int, insertFactor *int) {
	if root == nil {
		root = &TreeNode{data: data}
		return
	}

	if root.left == nil {
		root.left = &TreeNode{data: data}
		*insertFactor++
		return
	} else if root.right == nil {
		root.right = &TreeNode{data: data}
		*insertFactor++
		return
	}

	if *insertFactor%2 == 0 {
		t.Insert(root.left, data, insertFactor)
	} else {
		t.Insert(root.right, data, insertFactor)
	}
}

func (t *TreeNode) FindNode(root *TreeNode, data int) bool {
	if root == nil {
		return false
	}

	if root.data == data {
		return true
	}
	return t.FindNode(root.left, data) || t.FindNode(root.right, data)
}

func (t *TreeNode) Delete(root *TreeNode, data int) bool {
	if root == nil {
		return false
	}

	var nodeToDelete, deepestNode, parentDeepestNode *TreeNode
	q := []*TreeNode{root}
	for len(q) != 0 {
		deepestNode = q[0]
		q = q[1:]

		if deepestNode.data == data {
			nodeToDelete = deepestNode
		}
		if deepestNode.left != nil {
			parentDeepestNode = deepestNode
			if deepestNode.left.data == data {
				nodeToDelete = deepestNode.left
			}
			q = append(q, deepestNode.left)
		}

		if deepestNode.right != nil {
			parentDeepestNode = deepestNode
			if deepestNode.right.data == data {
				nodeToDelete = deepestNode.right
			}
			q = append(q, deepestNode.right)
		}

	}

	if nodeToDelete == nil {
		return false
	}

	if nodeToDelete == deepestNode {
		if parentDeepestNode.left == deepestNode {
			parentDeepestNode.left = nil
		} else if parentDeepestNode.right == deepestNode {
			parentDeepestNode.right = nil
		}
		return true
	}

	nodeToDelete.data = deepestNode.data
	if parentDeepestNode.left == deepestNode {
		parentDeepestNode.left = nil
	} else if parentDeepestNode.right == deepestNode {
		parentDeepestNode.right = nil
	}

	return true
}

func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	inOrderTraversal(root.left)
	fmt.Print(root.data, " ")
	inOrderTraversal(root.right)
}

func main() {
	root := &TreeNode{data: 10}
	insertFactor := 0

	root.Insert(root, 5, &insertFactor)
	root.Insert(root, 15, &insertFactor)
	root.Insert(root, 3, &insertFactor)
	root.Insert(root, 7, &insertFactor)
	root.Insert(root, 12, &insertFactor)
	root.Insert(root, 17, &insertFactor)

	fmt.Println("Tree before deletion:")
	inOrderTraversal(root)
	fmt.Println()

	root.Delete(root, 5)

	fmt.Println("Tree after deletion:")
	inOrderTraversal(root)
	fmt.Println()
}
