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

func TestGraph_MST(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")

	graph.AddEdge("0", "1", 5)
	graph.AddEdge("1", "0", 5)

	graph.AddEdge("0", "5", 3)
	graph.AddEdge("5", "0", 3)

	graph.AddEdge("1", "2", 10)
	graph.AddEdge("2", "1", 10)

	graph.AddEdge("1", "4", 1)
	graph.AddEdge("4", "1", 1)

	graph.AddEdge("1", "6", 4)
	graph.AddEdge("6", "1", 4)

	graph.AddEdge("2", "3", 5)
	graph.AddEdge("3", "2", 5)

	graph.AddEdge("2", "6", 8)
	graph.AddEdge("6", "2", 8)

	graph.AddEdge("3", "4", 7)
	graph.AddEdge("4", "3", 7)

	graph.AddEdge("3", "6", 9)
	graph.AddEdge("6", "3", 9)

	graph.AddEdge("4", "5", 6)
	graph.AddEdge("5", "4", 6)

	graph.AddEdge("4", "6", 2)
	graph.AddEdge("6", "4", 2)

	//	graph.PrintGraph()

	// expect MST: [{1 4 1} {4 6 2} {0 5 3} {0 1 5} {2 3 5} {4 3 7}]
	fmt.Printf("KruskalMST: %v \n", graph.KruskalMST())

	// expect MST: [{1 4 1} {4 6 2} {0 5 3} {0 1 5} {2 3 5} {4 3 7}]
	fmt.Printf("PrimMST: \n")
	for _, edge := range graph.PrimMST("2") {
		fmt.Printf("From: %v, To: %v, Weight: %v\n", edge.From, edge.To, edge.Weight)
	}
}

func TestGraph_ShortPath(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")
	graph.AddVertex("6")

	graph.AddEdge("0", "1", 5)
	graph.AddEdge("1", "0", 5)

	graph.AddEdge("0", "5", 3)
	graph.AddEdge("5", "0", 3)

	graph.AddEdge("1", "2", 10)
	graph.AddEdge("2", "1", 10)

	graph.AddEdge("1", "4", 1)
	graph.AddEdge("4", "1", 1)

	graph.AddEdge("1", "6", 4)
	graph.AddEdge("6", "1", 4)

	graph.AddEdge("2", "3", 5)
	graph.AddEdge("3", "2", 5)

	graph.AddEdge("2", "6", 8)
	graph.AddEdge("6", "2", 8)

	graph.AddEdge("3", "4", 7)
	graph.AddEdge("4", "3", 7)

	graph.AddEdge("3", "6", 9)
	graph.AddEdge("6", "3", 9)

	graph.AddEdge("4", "5", 6)
	graph.AddEdge("5", "4", 6)

	graph.AddEdge("4", "6", 2)
	graph.AddEdge("6", "4", 2)

	//	graph.PrintGraph()

	// expect ShortPath: map[0:15 1:10 2:0 3:5 4:10 5:16 6:8]
	shortPath, prevVertex := graph.DijkstraShortPath("2")
	fmt.Printf("DijkstraShortPath: %v \n", shortPath)

	// expect 2 to 5 short path is:  [2 6 4 5]
	path := ReconstructPath(prevVertex, "2", "5")
	fmt.Println("2 to 5 short path is: ", path)

	// expect ShortPath: map[0:15 1:10 2:0 3:5 4:10 5:16 6:8]
	shortPath, prevVertex2, hasNegativeCycle := graph.BellmanFordShortPath("2")
	fmt.Printf("BellmanFordShortPath: %v, hasNegativeCycle: %v \n", shortPath, hasNegativeCycle)

	// expect 2 to 5 short path is:  [2 6 4 5]
	path2 := ReconstructPath(prevVertex2, "2", "5")
	fmt.Println("2 to 5 short path is: ", path2)
}

func TestGraph_NegativeCycleShortPath(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")
	graph.AddVertex("4")
	graph.AddVertex("5")

	graph.AddEdge("0", "1", 5)

	graph.AddEdge("1", "2", 6)
	graph.AddEdge("1", "4", -4)

	graph.AddEdge("2", "4", -3)
	graph.AddEdge("2", "5", -2)

	graph.AddEdge("3", "2", 4)

	graph.AddEdge("4", "3", 1)
	graph.AddEdge("4", "5", 6)

	graph.AddEdge("5", "0", 3)
	graph.AddEdge("5", "1", 7)

	//	graph.PrintGraph()

	// expect BellmanFordShortPath: map[0:0 1:5 2:6 3:2 4:1 5:4], hasNegativeCycle: false
	shortPath, prevVertex, hasNegativeCycle := graph.BellmanFordShortPath("0")
	fmt.Printf("BellmanFordShortPath: %v, hasNegativeCycle: %v \n", shortPath, hasNegativeCycle)

	// expect 0 to 5 short path is:  [0 1 4 3 2 5]
	path := ReconstructPath(prevVertex, "0", "5")
	fmt.Println("0 to 5 short path is: ", path)
}

func TestGraph_FloydWarshall_ShortPath(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("0")
	graph.AddVertex("1")
	graph.AddVertex("2")
	graph.AddVertex("3")

	graph.AddEdge("0", "1", 2)
	graph.AddEdge("0", "2", 6)
	graph.AddEdge("0", "3", 8)

	graph.AddEdge("1", "2", -2)
	graph.AddEdge("1", "3", 3)

	graph.AddEdge("2", "0", 4)
	graph.AddEdge("2", "3", 1)

	//graph.PrintGraph()

	// expect BellmanFordShortPath: map[0:0 1:5 2:6 3:2 4:1 5:4], hasNegativeCycle: false
	shortPath, hasNegativeCycle := graph.FloydWarshallShortPath()

	if hasNegativeCycle {
		fmt.Println("圖中存在負權重循環")
	} else {
		fmt.Println("所有點對最短距離：")
		for from, row := range shortPath {
			for to, d := range row {
				fmt.Printf("%s → %s: %d\n", from, to, d)
			}
		}
	}

}

func graphToSlice(graph *Graph) []string {
	var result []string

	for _, vertex := range graph.GetVertices() {
		result = append(result, vertex.Id)
	}

	return result
}

func ReconstructPath(prev map[string]string, start, end string) []string {
	var path []string
	for to := end; to != ""; to = prev[to] {
		path = append([]string{to}, path...)
	}

	if path[0] != start {
		return []string{} // can not arrive
	}
	return path
}
