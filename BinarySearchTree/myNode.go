package BinarySearchTree

import (
	"golang.org/x/exp/constraints" // 用於泛型的比較約束
)

type Node[T constraints.Ordered] struct {
	Data T
	Left *Node[T]
	Right *Node[T]
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{
		Data: data,
		Left: nil,
		Right: nil,
	}
}