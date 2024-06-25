package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vietpham102301/go-dsa-practice/algorithms/cracking_the_coding_interviews/queue"
	"github.com/vietpham102301/go-dsa-practice/algorithms/cracking_the_coding_interviews/stack"
)

const MAX = 100

var V, E int

var graph [MAX][]int

var visited [MAX]bool
var path [MAX]int

func DFS(src int) {
	for i := 0; i < V; i++ {
		path[i] = -1
	}

	s := stack.Stack[int]{}

	visited[src] = true
	s.Push(src)

	for !s.IsEmpty() {
		u, _ := s.Pop()
		for i := 0; i < len(graph[*u]); i++ {
			v := graph[*u][i]
			if !visited[v] {
				visited[v] = true
				s.Push(v)
				path[v] = *u
			}
		}
	}

}

func BFS(src int) {
	for i := 0; i < V; i++ {
		path[i] = -1
	}

	q := queue.Queue[int]{}
	visited[src] = true
	q.Add(src)

	for !q.IsEmpty() {
		u, _ := q.Remove()
		for i := 0; i < len(graph[*u]); i++ {
			vertexToVisit := graph[*u][i]
			if !visited[vertexToVisit] {
				visited[vertexToVisit] = true
				q.Add(vertexToVisit)
				path[vertexToVisit] = *u
			}
		}
	}
}

func printPath(d int) {
	myStack := stack.Stack[int]{}
	for d != -1 {
		myStack.Push(d)
		d = path[d]
	}

	for !myStack.IsEmpty() {
		vertex, _ := myStack.Pop()
		if myStack.Head != nil {
			fmt.Printf("%d->", *vertex)
		} else {
			fmt.Printf("%d", *vertex)
		}
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	V, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	E, _ = strconv.Atoi(scanner.Text())

	fmt.Println(V, E)

	var u, v int

	for i := 0; i < E; i++ {
		scanner.Scan()
		u, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ = strconv.Atoi(scanner.Text())
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	DFS(0)
	printPath(5)
}
