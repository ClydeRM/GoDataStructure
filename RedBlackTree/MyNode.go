package RedBlackTree

import (
	"golang.org/x/exp/constraints" // 用於泛型的比較約束
)

type Color bool

const (
	RED Color = false
	BLACK Color = true
)

type Node[T constraints.Ordered] struct {
	Data   T
	Color Color
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

func NewNode[T constraints.Ordered](data T, color Color) *Node[T] {
	return &Node[T]{
		Data:   data,
		Color: color,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}
