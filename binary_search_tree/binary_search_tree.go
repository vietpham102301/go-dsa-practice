package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

var count int

func (n *Node) Insert(k int) {
	if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	} else if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.Key > k {
		return n.Left.Search(k)
	} else if n.Key < k {
		return n.Right.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 10}

	tree.Insert(20)
	tree.Insert(5)

	tree.Insert(30)
	tree.Insert(71)
	tree.Insert(39)
	tree.Insert(53)
	tree.Insert(14)

	fmt.Println(tree)

	fmt.Println(tree.Search(53))
	fmt.Println(count)

}
