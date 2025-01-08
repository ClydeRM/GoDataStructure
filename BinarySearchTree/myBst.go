package BinarySearchTree

import "golang.org/x/exp/constraints"

type BST[T constraints.Ordered] struct {
	root *Node[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

type Tree[T constraints.Ordered] interface {
	Insert(data T) *Node[T]
	Search(data T) *Node[T]
	Delete(data T) *Node[T]
	InOrderTraversal() []T
	PreOrderTraversal() []T
	PostOrderTraversal() []T
	Hight() int
}