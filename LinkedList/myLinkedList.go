package LinkedList

import "fmt"

type Node struct {
	Data int
	next *Node // next Node mem address
}

type LinkedList struct {
	head *Node // head Node pointer
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head // tem save current head as second
	l.head = n // asign new head node
	l.head.next = second // point second node
	l.length++
}
func PrintLinkedList() {
	fmt.Println("Hello Linked list.")
}