package BinarySearchTree

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	type args struct {
		dataList []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// test cases
		{
			"Prepend_EmptyTree",
			args{
				dataList: []int{},
			},
			[]int{},
		},
		{
			"Prepend_SingleNode",
			args{
				dataList: []int{1},
			},
			[]int{1},
		},
		{
			"Prepend_MultipleNode",
			args{
				dataList: []int{1,2,3},
			},
			[]int{1,2,3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST[int]()

			for _, data := range tt.args.dataList {
				tree.Insert(data)
			}
			gotValues := inOrderTraversal(tree.root)

			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("BST : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

// 輔助函數：中序遍歷樹並返回節點資料的切片
func inOrderTraversal[T constraints.Ordered](node *Node[T]) []T {
	if node == nil {
		return []T{}
	}
	result := []T{}
	result = append(result, inOrderTraversal(node.Left)...)
	result = append(result, node.Data)
	result = append(result, inOrderTraversal(node.Right)...)
	return result
}