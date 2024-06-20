package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (n *Node) inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	n.inOrderTraversal(node.left)
	fmt.Println("Visited Node: ", node.data)
	n.inOrderTraversal(node.right)
}

func (n *Node) preOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Println("Visited Node: ", node.data)
	n.preOrderTraversal(node.left)
	n.preOrderTraversal(node.right)
}

func (n *Node) postOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	n.postOrderTraversal(node.left)
	n.postOrderTraversal(node.right)
	fmt.Println("Visited Node: ", node.data)
}

func (n *Node) insert(val int) {
	if n.data > val {
		if n.left == nil {
			n.left = &Node{data: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: val}
		} else {
			n.right.insert(val)
		}
	}
}

func main() {
	tree := &Node{data: 5}

	tree.insert(3)
	tree.insert(7)
	tree.insert(2)
	tree.insert(4)
	tree.insert(6)
	tree.insert(8)
	tree.insert(1)

	fmt.Println("In-Order traversal")
	tree.inOrderTraversal(tree)

	fmt.Println("Pre-Order traversal")
	tree.preOrderTraversal(tree)

	fmt.Println("Post-Order traversal")
	tree.postOrderTraversal(tree)
}
