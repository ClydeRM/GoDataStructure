package BinarySearchTree

import (
	"fmt"
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
			"Insert_EmptyTree",
			args{
				dataList: []int{},
			},
			[]int{},
		},
		{
			"Insert_SingleNode",
			args{
				dataList: []int{1},
			},
			[]int{1},
		},
		{
			"Insert_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
			},
			[]int{1, 2, 3},
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

func TestBST_Search(t *testing.T) {
	type args struct {
		dataList []int
		target   int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		// test cases
		{
			"Search_EmptyTree",
			args{
				dataList: []int{},
				target:   0,
			},
			nil,
		},
		{
			"Search_SingleNode",
			args{
				dataList: []int{1},
				target:   1,
			},
			&Node[int]{Data: 1}, // 預期返回的節點
		},
		{
			"Search_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
				target:   1,
			},
			&Node[int]{Data: 1}, // 預期返回的節點
		},
		{
			"Search_NotFound",
			args{
				dataList: []int{1, 2, 3},
				target:   4,
			},
			nil, // 找不到時應返回 nil
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST[int]()

			// 插入節點
			for _, data := range tt.args.dataList {
				tree.Insert(data)
			}

			// 呼叫 Search 查找目標節點
			result := tree.Search(tt.args.target)

			// 驗證結果
			if (result == nil && tt.want != nil) || (result != nil && result.Data != tt.want.Data) {
				t.Errorf("Search(%d): got = %v, want = %v", tt.args.target, result, tt.want)
			}
		})
	}
}

func TestBST_PreOrder(t *testing.T) {
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
			"PreOrder_EmptyTree",
			args{
				dataList: []int{},
			},
			[]int{},
		},
		{
			"PreOrder_SingleNode",
			args{
				dataList: []int{1},
			},
			[]int{1},
		},
		{
			"PreOrder_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
			},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.PreOrderTraversal(bst.root)

			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("PreOrder : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_InOrder(t *testing.T) {
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
			"InOrder_EmptyTree",
			args{
				dataList: []int{},
			},
			[]int{},
		},
		{
			"InOrder_SingleNode",
			args{
				dataList: []int{1},
			},
			[]int{1},
		},
		{
			"InOrder_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
			},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.InOrderTraversal(bst.root)

			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("InOrder : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_PostOrder(t *testing.T) {
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
			"PostOrder_EmptyTree",
			args{
				dataList: []int{},
			},
			[]int{},
		},
		{
			"PostOrder_SingleNode",
			args{
				dataList: []int{1},
			},
			[]int{1},
		},
		{
			"PostOrder_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
			},
			[]int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.PostOrderTraversal(bst.root)

			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("PostOrder : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_Min(t *testing.T) {
	type args struct {
		dataList []int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		// test cases
		{
			"Min_EmptyTree",
			args{
				dataList: []int{},
			},
			nil,
		},
		{
			"Min_SingleNode",
			args{
				dataList: []int{1},
			},
			&Node[int]{Data: 1},
		},
		{
			"Min_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
			},
			&Node[int]{Data: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.Min(bst.root)

			if (gotValues == nil && tt.want != nil) || (gotValues != nil && gotValues.Data != tt.want.Data) {
				t.Errorf("Min : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_Max(t *testing.T) {
	type args struct {
		dataList []int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		// test cases
		{
			"Max_EmptyTree",
			args{
				dataList: []int{},
			},
			nil,
		},
		{
			"Max_SingleNode",
			args{
				dataList: []int{1},
			},
			&Node[int]{Data: 1},
		},
		{
			"Max_MultipleNode",
			args{
				dataList: []int{1, 2, 3},
			},
			&Node[int]{Data: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.Max(bst.root)

			if (gotValues == nil && tt.want != nil) || (gotValues != nil && gotValues.Data != tt.want.Data) {
				t.Errorf("Min : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_InOrderSuccessor(t *testing.T) {
	type args struct {
		dataList []int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		// test cases
		{
			"InOrderSuccessor_EmptyTree",
			args{
				dataList: []int{},
			},
			nil,
		},
		{
			"InOrderSuccessor_SingleNode",
			args{
				dataList: []int{1},
			},
			nil,
		},
		{
			"InOrderSuccessor_MultipleNode",
			args{
				dataList: []int{2, 1, 3},
			},
			&Node[int]{Data: 3},
		},
		{
			"InOrderSuccessor_NoRight",
			args{
				dataList: []int{2, 1},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.InOrderSuccessor(bst.root)

			if (gotValues == nil && tt.want != nil) || (gotValues != nil && gotValues.Data != tt.want.Data) {
				t.Errorf("InOrderSuccessor : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_InOrderPredecessor(t *testing.T) {
	type args struct {
		dataList []int
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		// test cases
		{
			"InOrderPredecessor_EmptyTree",
			args{
				dataList: []int{},
			},
			nil,
		},
		{
			"InOrderPredecessor_SingleNode",
			args{
				dataList: []int{1},
			},
			nil,
		},
		{
			"InOrderPredecessor_MultipleNode",
			args{
				dataList: []int{2, 1, 3},
			},
			&Node[int]{Data: 1},
		},
		{
			"InOrderPredecessor_NoLeft",
			args{
				dataList: []int{2, 3},
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, data := range tt.args.dataList {
				bst.Insert(data)
			}
			gotValues := bst.InOrderPredecessor(bst.root)

			if (gotValues == nil && tt.want != nil) || (gotValues != nil && gotValues.Data != tt.want.Data) {
				t.Errorf("InOrderPredecessor : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}

func TestBST_Delete(t *testing.T) {
	type args struct {
		dataList []int
		target   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// test cases
		{
			"Delete_EmptyTree",
			args{
				dataList: []int{},
				target:   0,
			},
			[]int{},
		},
		{
			"Delete_SingleNode",
			args{
				dataList: []int{10},
				target:   10,
			},
			[]int{},
		},
		{
			"Delete_MultipleNode_NotFound",
			args{
				dataList: []int{10, 5, 15, 12, 20},
				target:   9,
			},
			[]int{5, 10, 12, 15, 20},
		},
		{
			"Delete_MultipleNode_LeafNode_left",
			args{
				dataList: []int{10, 5, 15, 12, 20},
				target:   12,
			},
			[]int{5, 10, 15, 20},
		},
		{
			"Delete_MultipleNode_LeafNode_right",
			args{
				dataList: []int{10, 5, 15, 12, 20},
				target:   20,
			},
			[]int{5, 10, 12, 15},
		},
		{
			"Delete_MultipleNode_OneChild_left",
			args{
				dataList: []int{10, 5, 4, 15, 12, 20},
				target:   5,
			},
			[]int{4, 10, 12, 15, 20},
		},
		{
			"Delete_MultipleNode_OneChild_right",
			args{
				dataList: []int{10, 5, 6, 15, 12, 20},
				target:   6,
			},
			[]int{5, 10, 12, 15, 20},
		},
		{
			"Delete_MultipleNode_BothChildren",
			args{
				dataList: []int{10, 5, 15, 12, 20},
				target:   15,
			},
			[]int{5, 10, 12, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST[int]()

			for _, data := range tt.args.dataList {
				tree.Insert(data)
			}
			fmt.Printf("before: %v \n", inOrderTraversal(tree.root))
			tree.Delete(tt.args.target)

			gotValues := inOrderTraversal(tree.root)
			fmt.Printf("after: %v \n", gotValues)

			if !reflect.DeepEqual(gotValues, tt.want) {
				t.Errorf("BST : got = %v, want = %v", gotValues, tt.want)
			}
		})
	}
}


func TestBST_Height(t *testing.T) {
	type args struct {
		dataList []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test cases
		{
			"Height_EmptyTree",
			args{
				dataList: []int{},
			},
			-1,
		},
		{
			"Height_SingleNode",
			args{
				dataList: []int{10},
			},
			0,
		},
		{
			"Height_MultipleNode_BothChildren",
			args{
				dataList: []int{10, 5, 15, 12, 20},
			},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST[int]()

			for _, data := range tt.args.dataList {
				tree.Insert(data)
			}
			fmt.Printf("BST: %v \n", inOrderTraversal(tree.root))
			gotValues := tree.Height()

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
	var result []T
	result = append(result, inOrderTraversal(node.left)...)
	result = append(result, node.Data)
	result = append(result, inOrderTraversal(node.right)...)
	return result
}
