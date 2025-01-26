package RedBlackTree

import (
	//	"fmt"
	//	"golang.org/x/exp/constraints"
	//	"reflect"
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