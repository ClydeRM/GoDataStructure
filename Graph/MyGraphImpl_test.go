package Graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGraph_AddVertex(t *testing.T) {
	type args struct {
		dataList []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// test cases
		{
			"TestGraph_AddVertex_non_exist",
			args{
				dataList: []string{"A"},
			},
			[]string{"A"},
		},
		{
			"TestGraph_AddVertex_exist",
			args{
				dataList: []string{"A", "A"},
			},
			[]string{"A"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewGraph()

			for _, vertexId := range tt.args.dataList {
				graph.AddVertex(vertexId)
			}

			graph.PrintGraph()

			got := graphToSlice(graph)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph AddVertex failed: got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindStronglyConnectedComponents(t *testing.T) {
	graph := NewGraph()
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

	graph.AddEdge("1", "2", 1)
	graph.AddEdge("1", "4", 1)

	graph.AddEdge("2", "0", 1)
	graph.AddEdge("2", "3", 1)
	graph.AddEdge("2", "5", 1)

	graph.AddEdge("3", "2", 1)

	graph.AddEdge("4", "5", 1)
	graph.AddEdge("4", "6", 1)

	graph.AddEdge("5", "4", 1)
	graph.AddEdge("5", "6", 1)
	graph.AddEdge("5", "7", 1)

	graph.AddEdge("6", "7", 1)

	graph.AddEdge("7", "8", 1)

	graph.AddEdge("8", "6", 1)


	graph.PrintGraph()
	fmt.Printf("BFS()0 : %v \n", graph.BFS("0"))
	fmt.Printf("BFS()4 : %v \n", graph.BFS("4"))
	fmt.Printf("BFS()6 : %v \n", graph.BFS("6"))

	// expect Strong Connected Component : [[0 2 1 3] [5 4] [6 8 7]]
	fmt.Printf("Strong Connected Component : %v \n", graph.FindStronglyConnectedComponents())
}

func graphToSlice(graph *Graph) []string {
	var result []string

	for _, vertex := range graph.GetVertices() {
		result = append(result, vertex.Id)
	}

	return result
}
