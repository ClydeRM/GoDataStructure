package main

import (
	"GoDataStructure/Graph"
	"fmt"
)

func main() {
	fmt.Println("run main.go..")
	graph := Graph.NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")

	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "D", 1)
	graph.AddEdge("B", "C", 1)
	graph.AddEdge("B", "E", 1)
	graph.AddEdge("D", "E", 1)

	graph.PrintGraph()
	result := graph.DFS("A")
	fmt.Println(result) // [A D E B C]

	result2 := graph.BFS("A")
	fmt.Println(result2) // [A B D C E]

}
