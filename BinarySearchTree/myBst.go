package BinarySearchTree

import "golang.org/x/exp/constraints"

type BST[T constraints.Ordered] struct {
	root *Node[T]
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

type Tree[T constraints.Ordered] interface {
	Root() *Node[T]
	IsEmpty() bool
	Insert(data T) *Node[T]
	Search(data T) *Node[T]
	Delete(data T)
	Min(root *Node[T]) *Node[T]
	Max(root *Node[T]) *Node[T]
	PreOrderTraversal(root *Node[T]) []T
	InOrderTraversal(root *Node[T]) []T
	PostOrderTraversal(root *Node[T]) []T
	InOrderSuccessor(node *Node[T]) *Node[T]   // InOrder visit: next visit node
	InOrderPredecessor(node *Node[T]) *Node[T] // InOrder visit: previous visit node

	Height() int
	Size() int
}
