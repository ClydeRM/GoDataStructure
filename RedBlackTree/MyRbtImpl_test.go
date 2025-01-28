package RedBlackTree

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"testing"
)

func TestRBT_Insert(t *testing.T) {
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
				dataList: []int{3, 1, 5, 2, 4},
			},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rbt := NewRBT[int]()

			for _, data := range tt.args.dataList {
				rbt.Insert(data)
			}

			// validate logic
			// 1. inorder traversal correctly
			got := rbt.InOrderTraversal(rbt.root)
			fmt.Printf("RBT: %v \n", got)
			if !isEqual(got, tt.want) {
				t.Errorf("RBT Insert failed: got %v, want %v", got, tt.want)
			}

			// 2. Complies with RBT standards
			if !isValidRBT(rbt.root, rbt.nilNode) {
				t.Errorf("RBT properties violated after Insert")
			}
		})
	}
}

func TestRBT_Delete_Iterative(t *testing.T) {
	type args struct {
		dataList []int
		target   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
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
				dataList: []int{1},
				target:   1,
			},
			[]int{},
		},
		{
			"Delete_RootNodeWithOneChild",
			args{
				dataList: []int{2, 1},
				target:   2,
			},
			[]int{1},
		},
		{
			"Delete_RootNodeWithTwoChildren",
			args{
				dataList: []int{10, 5, 15},
				target:   10,
			},
			[]int{5, 15},
		},
		{
			"Delete_NodeWithTwoChildren",
			args{
				dataList: []int{20, 10, 30, 5, 15},
				target:   10,
			},
			[]int{5, 15, 20, 30},
		},
		{
			"Delete_NodeWithOneChild",
			args{
				dataList: []int{8, 4, 10, 2},
				target:   4,
			},
			[]int{2, 8, 10},
		},
		{
			"Delete_LeafNode",
			args{
				dataList: []int{8, 4, 10, 2},
				target:   2,
			},
			[]int{4, 8, 10},
		},
		{
			"Delete_MultipleNodes",
			args{
				dataList: []int{40, 20, 60, 10, 30, 50, 70},
				target:   20,
			},
			[]int{10, 30, 40, 50, 60, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 初始化紅黑樹
			rbt := NewRBT[int]()
			for _, data := range tt.args.dataList {
				rbt.Insert(data)
			}
			fmt.Printf("before: %v \n", rbt.InOrderTraversal(rbt.root))
			fmt.Printf("RBT Delete: %v \n", tt.args.target)
			rbt.Delete(tt.args.target)

			// 驗證樹中序遍歷結果是否正確
			got := rbt.InOrderTraversal(rbt.root)
			fmt.Printf("after: %v \n", got)

			if !isEqual(got, tt.want) {
				t.Errorf("RBT Delete failed: got %v, want %v", got, tt.want)
			}

			// 驗證紅黑樹是否依然符合性質
			if !isValidRBT(rbt.root, rbt.nilNode) {
				t.Errorf("RBT properties violated after Delete")
			}
		})
	}
}

func isEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 檢查紅黑樹是否符合性質
func isValidRBT[T constraints.Ordered](node, nilNode *Node[T]) bool {
	if node == nilNode {
		return true
	}

	// 1. 檢查根節點是否為黑色
	if node.parent == nil && node.Color != BLACK {
		return false
	}

	// 2. 檢查沒有兩個連續的紅色節點
	if node.Color == RED {
		if node.left != nilNode && node.left.Color == RED {
			return false
		}
		if node.right != nilNode && node.right.Color == RED {
			return false
		}
	}

	// 3. 檢查黑色高度是否一致
	leftBlackHeight := countBlackHeight(node.left, nilNode)
	rightBlackHeight := countBlackHeight(node.right, nilNode)
	if leftBlackHeight != rightBlackHeight {
		return false
	}

	// 遞迴檢查子樹
	return isValidRBT(node.left, nilNode) && isValidRBT(node.right, nilNode)
}

// 計算黑色高度
func countBlackHeight[T constraints.Ordered](node, nilNode *Node[T]) int {
	if node == nilNode {
		return 1
	}
	leftHeight := countBlackHeight(node.left, nilNode)
	rightHeight := countBlackHeight(node.right, nilNode)
	if node.Color == BLACK {
		return _max(leftHeight, rightHeight) + 1
	}
	return _max(leftHeight, rightHeight)
}

// 返回兩個數中的較大值
func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
