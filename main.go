package main

import (
	"GoDataStructure/LinkedList"
	"fmt"
)

func main() {
	fmt.Println("run main.go..")
	mylist := LinkedList.LinkedList[int]{}
	nodeA := &LinkedList.Node[int]{Data: 10}
	nodeB := &LinkedList.Node[int]{Data: 20}
	nodeC := &LinkedList.Node[int]{Data: 30}
	nodeD := &LinkedList.Node[int]{Data: 40}
	nodeE := &LinkedList.Node[int]{Data: 50}

	mylist.Prepend(nodeE)
	mylist.Prepend(nodeD)
	mylist.Prepend(nodeC)
	mylist.Prepend(nodeB)
	mylist.Prepend(nodeA)

	mylist.PrintListData()
	fmt.Println(nodeB)

	//	mylist.Update(100,5)
	mylist.Reverse()

	mylist.PrintListData()
	fmt.Println(nodeB)

}
