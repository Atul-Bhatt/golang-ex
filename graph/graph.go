package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("One", "Two")
	addEdge("One", "Three")
	addEdge("Three", "Two")
	addEdge("Three", "One")
	fmt.Println(hasEdge("One", "Two"))
	fmt.Println(hasEdge("Three", "Four"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
