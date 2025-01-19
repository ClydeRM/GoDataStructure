package BinarySearchTree

import (
	"golang.org/x/exp/constraints" // 用於泛型的比較約束
)

type Node[T constraints.Ordered] struct {
	Data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{
		Data:   data,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}
