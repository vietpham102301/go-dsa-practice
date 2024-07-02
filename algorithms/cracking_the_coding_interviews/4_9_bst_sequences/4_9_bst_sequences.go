package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

func allSequences(node *TreeNode) []*list.List {
	result := make([]*list.List, 0)

	if node == nil {
		result = append(result, list.New())
		return result
	}

	prefix := list.New()
	prefix.PushFront(node.data)

	// recurse on left and right subtrees
	leftSeq := allSequences(node.left)
	rightSeq := allSequences(node.right)

	// weave together each list from the left and right sides
	for _, left := range leftSeq {
		for _, right := range rightSeq {
			weaved := make([]*list.List, 0)
			weaveLists(left, right, &weaved, prefix)
			result = append(result, weaved...)
		}
	}

	return result
}

func weaveLists(first *list.List, second *list.List, results *[]*list.List, prefix *list.List) {
	if first.Len() == 0 || second.Len() == 0 {
		result := list.New()
		for e := prefix.Front(); e != nil; e = e.Next() {
			result.PushBack(e.Value)
		}
		result.PushBackList(first)
		result.PushBackList(second)

		*results = append(*results, result)
		return
	}

	headFirst := first.Remove(first.Front())
	prefix.PushBack(headFirst)
	weaveLists(first, second, results, prefix)
	prefix.Remove(prefix.Back())
	first.PushFront(headFirst)

	headSecond := second.Remove(second.Front())
	prefix.PushBack(headSecond)
	weaveLists(first, second, results, prefix)
	prefix.Remove(prefix.Back())
	second.PushFront(headSecond)
}

func main() {
	root := &TreeNode{data: 9}
	root.left = &TreeNode{data: 3}
	root.right = &TreeNode{data: 15}

	root.left.left = &TreeNode{data: 1}
	root.left.right = &TreeNode{data: 5}

	root.right.right = &TreeNode{data: 20}

	res := allSequences(root)
	count := 0
	for _, l := range res {
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d, ", e.Value)
		}
		fmt.Println()
		count++
	}
	fmt.Println(count)
}
