package main

import (
	"GoDataStructure/Graph"
	"fmt"
)

func main() {
	fmt.Println("run main.go..")
	graph := Graph.NewGraph()
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")
	graph.AddVertex("7")
	graph.AddVertex("8")

	graph.AddEdge("0", "1", 1)
	graph.AddEdge("1", "4", 1)
	graph.AddEdge("1", "5", 1)
	graph.AddEdge("4", "5", 1)
	graph.AddEdge("5", "7", 1)

	graph.AddEdge("3", "6", 1)
	graph.AddEdge("6", "8", 1)

	graph.PrintGraph()
	fmt.Printf("BFS()1 : %v \n", graph.BFS("0"))
	fmt.Printf("BFS()2 : %v \n", graph.BFS("2"))
	fmt.Printf("BFS()3 : %v \n", graph.BFS("3"))

	fmt.Printf("Component : %v \n", graph.FindConnectedComponents())
}
