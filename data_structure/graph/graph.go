package main

import "fmt"

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex //Linked List => comparing bigOh for operations
}

// Add vertex
func (g *Graph) addVertex(key int) {
	if !contains(g.vertices, key) {
		g.vertices = append(g.vertices, &Vertex{key: key, adjacent: []*Vertex{}})
	}
}

// Add edge
func (g *Graph) addEdge(from int, to int) {
	//checking if the "from" and "to" is already in the graph?
	if !contains(g.vertices, from) || !contains(g.vertices, to) {
		return
	}
	fromVertexIndex := getVertex(g.vertices, from)
	toVertexIndex := getVertex(g.vertices, to)
	fromVertex := g.vertices[fromVertexIndex]
	fromVertex.adjacent = append(fromVertex.adjacent, g.vertices[toVertexIndex])
}

// Get vertex return index of the vertex
func getVertex(vertices []*Vertex, key int) int {
	for i, v := range vertices {
		if v.key == key {
			return i
		}
	}
	return -1
}

// Contains
func contains(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if key == v.key {
			return true
		}
	}
	return false
}

func (g Graph) printGraph() {
	for _, i := range g.vertices {
		fmt.Printf("Vertex: %d -> ", i.key)
		for _, j := range i.adjacent {
			fmt.Printf("%d, ", j.key)
		}
		fmt.Println()
	}
}

func main() {
	graph := &Graph{}
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)
	graph.addVertex(4)

	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 3)
	graph.addEdge(3, 4)
	graph.addEdge(4, 1)

	graph.printGraph()
}
