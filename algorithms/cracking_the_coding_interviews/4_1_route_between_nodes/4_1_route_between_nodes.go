package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/vietpham102301/go-dsa-practice/algorithms/cracking_the_coding_interviews/queue"
)

const MAX = 100

var V, E int

var graph [MAX][]int

var path [MAX]int
var visited [MAX]bool

func Initialize() {
	for i := 0; i < V; i++ {
		path[i] = -1
	}
}

func ResetVisited() {
	for i := 0; i < V; i++ {
		visited[i] = false
	}
}

func BFS(src int) {
	ResetVisited()

	visited[src] = true
	q := queue.Queue[int]{}

	q.Add(src)

	for !q.IsEmpty() {
		u, _ := q.Remove()
		for i := 0; i < len(graph[*u]); i++ {
			vertexToVisit := graph[*u][i]
			if !visited[vertexToVisit] {
				q.Add(vertexToVisit)
				visited[vertexToVisit] = true
				path[vertexToVisit] = *u
			}
		}
	}
}

func IsHaveRoute(u, v int) bool {
	BFS(u)
	finnalNodeFromV := -1
	for v != -1 {
		if v != -1 {
			finnalNodeFromV = v
		}
		v = path[v]
	}

	return finnalNodeFromV == u
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	V, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	E, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < E; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())

		graph[u] = append(graph[u], v)
	}

	Initialize()

	fmt.Println(IsHaveRoute(1, 5))

}
