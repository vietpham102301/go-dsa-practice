package main

import (
	"fmt"

	"github.com/vietpham102301/go-dsa-practice/algorithms/cracking_the_coding_interviews/stack"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key           int
	adjacencyList []*Vertex
}

func (g *Graph) addVertex(key int) {
	if !isContains(g.vertices, key) {
		g.vertices = append(g.vertices, &Vertex{key: key, adjacencyList: []*Vertex{}})
	}
}

func (g *Graph) addEdge(from, to int) {
	var indexes []int
	for i, v := range g.vertices {
		if v.key == from || v.key == to {
			indexes = append(indexes, i)
		}
	}

	if len(indexes) != 2 {
		return
	}

	vertex1 := g.vertices[indexes[0]]
	vertex2 := g.vertices[indexes[1]]

	vertex2.adjacencyList = append(vertex2.adjacencyList, vertex1)
	vertex1.adjacencyList = append(vertex1.adjacencyList, vertex2)
}

func isContains(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.key == key {
			return true
		}
	}

	return false
}

func (g Graph) printGraph() {
	for _, v := range g.vertices {
		fmt.Printf("vertex: %d ->", v.key)
		for _, a := range v.adjacencyList {
			fmt.Printf("%d, ", a.key)
		}
		fmt.Println()
	}
}

// functionalize the DFS with stack and recursion

func dfsWithStack(graph Graph, startVertex int, des int) {
	var visited []int
	var path []int
	for i := 0; i < len(graph.vertices)+1; i++ {
		visited = append(visited, 0)
		path = append(path, -1)
	}

	myStack := stack.Stack[int]{}

	visited[startVertex] = 1
	myStack.Push(startVertex)

	for !myStack.IsEmpty() {
		visitedVertex, success := myStack.Pop()
		if success {
			adjList := graph.vertices[*visitedVertex-1].adjacencyList
			for j := 0; j < len(adjList); j++ {
				adjVertex := adjList[j].key
				if visited[adjVertex] != 1 {
					visited[adjVertex] = 1
					myStack.Push(adjVertex)
					if path[adjVertex] == -1 {
						path[adjVertex] = *visitedVertex
					}
				}
			}
		}
	}

	for des != -1 {
		myStack.Push(des)
		des = path[des]
	}

	for !myStack.IsEmpty() {
		vertex, _ := myStack.Pop()
		if myStack.Head == nil {
			fmt.Printf("%d", *vertex)
		} else {
			fmt.Printf("%d ->", *vertex)
		}
	}
}

// func main() {
// 	graph := Graph{}

// 	graph.addVertex(1)
// 	graph.addVertex(2)
// 	graph.addVertex(3)
// 	graph.addVertex(4)
// 	graph.addVertex(5)

// 	graph.addEdge(1, 2)
// 	graph.addEdge(1, 3)
// 	graph.addEdge(2, 3)
// 	graph.addEdge(3, 4)
// 	graph.addEdge(3, 5)
// 	graph.addEdge(4, 5)

// 	// graph.printGraph()
// 	startVertex := graph.vertices[0].key
// 	des := graph.vertices[4].key
// 	dfsWithStack(graph, startVertex, des)
// }
