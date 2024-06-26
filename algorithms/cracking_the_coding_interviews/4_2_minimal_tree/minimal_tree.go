package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int
}

func createMinimalTreeFromArr(arr []int) *Node {
	rootIndex := len(arr) / 2
	if len(arr) == 0 {
		return nil
	}

	rootNode := &Node{data: arr[rootIndex]}

	left := createMinimalTreeFromArr(arr[:rootIndex])

	right := createMinimalTreeFromArr(arr[rootIndex+1:])

	rootNode.left = left
	rootNode.right = right

	return rootNode
}

func inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.left)
	fmt.Printf("%d, ", node.data)
	inOrderTraversal(node.right)
}

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := createMinimalTreeFromArr(sortedArr)

	inOrderTraversal(root)
}
