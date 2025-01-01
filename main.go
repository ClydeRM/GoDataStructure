package main

import (
	"GoDataStructure/LinkedList"
	"fmt"
)

func main() {
	fmt.Println("Hello main.")
	LinkedList.PrintLinkedList()

	mylist := LinkedList.LinkedList{}
	nodeA := &LinkedList.Node{Data: 10}
	nodeB := &LinkedList.Node{Data: 20}
	nodeC := &LinkedList.Node{Data: 30}
	nodeD := &LinkedList.Node{Data: 40}

	mylist.Prepend(nodeD)
	mylist.Prepend(nodeB)
	mylist.Prepend(nodeC)
	mylist.Prepend(nodeA)
}
