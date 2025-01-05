package LinkedList

type Node[T any] struct {
	Data T
	prev *Node[T]
	next *Node[T]
}