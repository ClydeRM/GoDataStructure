package Graph

import (
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

func graphToSlice(graph *Graph) []string {
	var result []string

	for _, vertex := range graph.GetVertices() {
		result = append(result, vertex.Id)
	}

	return result
}
