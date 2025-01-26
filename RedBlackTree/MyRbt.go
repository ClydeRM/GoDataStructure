package RedBlackTree

import "golang.org/x/exp/constraints"

type RBT[T constraints.Ordered] struct {
    root    *Node[T]
    size    int
    nilNode *Node[T] // 哨兵節點 (表示 NIL)
}

func NewRBT[T constraints.Ordered]() *RBT[T] {
    return &RBT[T]{nil, 0, nil}
}

type Tree[T constraints.Ordered] interface {
    Insert(data T)      // 插入節點
    Delete(data T)      // 刪除節點
    Search(data T) bool // 查找某個值是否存在
    Min() (T, bool)     // 找到最小值，若樹為空則返回 false
    Max() (T, bool)     // 找到最大值，若樹為空則返回 false
    Size() int          // 返回紅黑樹中節點的數量
    Height() int        // 返回紅黑樹的高度
    IsEmpty() bool      // 判斷紅黑樹是否為空
    InOrderTraversal(root *Node[T]) []T
}
