package LinkedList

import "fmt"

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		nil, nil, 0,
	}
}

func (l *LinkedList[T]) Prepend(newNode *Node[T]) {
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		second := l.head
		newNode.next = second
		second.prev = newNode
		l.head = newNode
	}

	l.length++
}

func (l *LinkedList[T]) Append(newNode *Node[T]) {
	if l.tail != nil {
		newNode.prev = l.tail
		l.tail.next = newNode
	} else {
		l.head = newNode
	}

	l.tail = newNode
	l.length++
}

func (l *LinkedList[T]) Pop() {
	if l.IsEmpty(){
		fmt.Printf("Empty list: %v\n", l)
		return
	}

	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return
	}

	l.tail = l.tail.prev
	l.tail.next = nil
	l.length--
}

func (l *LinkedList[T]) Insert(newNode *Node[T], index int) {
	if index < 0 || index > l.length {
		fmt.Printf("invalid action index: %v\n", index)
		return
	}

	if index == 0 {
		l.Prepend(newNode)
		return
	}

	if index == l.length {
		l.Append(newNode)
		return
	}

	currNode := l.head

	for prevIndex := index - 1; prevIndex > 0; prevIndex-- {
		currNode = currNode.next
	}

	originNext := currNode.next
	newNode.next = originNext
	newNode.prev = currNode

	currNode.next = newNode
	originNext.prev = newNode

	l.length++
}

func (l *LinkedList[T]) Remove(index int) {
	if index < 0 || index > l.length {
		fmt.Printf("invalid action index: %v\n", index)
		return
	}

	if isEmpty := l.IsEmpty(); isEmpty{
		fmt.Printf("Empty list: %v\n", l)
		return
	}

	if index == 0 {
		currHead := l.head
		newHead := l.head.next
		newHead.prev = nil
		currHead.next = nil
		l.head = newHead
		l.length--
		return
	}

	if index == l.length {
		l.Pop()
		return
	}

	currNode := l.head

	for prevIndex := index - 1; prevIndex > 0; prevIndex-- {
		currNode = currNode.next
	}

	removeTarget := currNode.next
	newNext := currNode.next.next

	removeTarget.prev = nil
	removeTarget.next = nil

	currNode.next = newNext
	newNext.prev = currNode

	l.length--
}

func (l *LinkedList[T]) Get(index int) *Node[T]{
	if index < 0 || index > l.length {
		fmt.Printf("invalid action index: %v\n", index)
		return nil
	}

	if isEmpty := l.IsEmpty(); isEmpty{
		fmt.Printf("Empty list: %v\n", l)
		return nil
	}

	if index == 0 {
		return l.head
	}

	if index == l.length {
		return l.tail
	}

	currNode := l.head

	for ; index > 0; index-- {
		currNode = currNode.next
	}

	return currNode
}

func (l *LinkedList[T]) Update(newData T, index int) {
	if index < 0 || index > l.length {
		fmt.Printf("invalid action index: %v\n", index)
		return
	}

	if isEmpty := l.IsEmpty(); isEmpty{
		fmt.Printf("Empty list: %v\n", l)
	}

	if index == 0 {
		l.head.Data = newData
		return
	}

	if index == l.length {
		l.tail.Data = newData
		return
	}

	currNode := l.head

	for ; index > 0; index-- {
		currNode = currNode.next
	}

	currNode.Data = newData
}
func (l *LinkedList[T]) Reverse() {
	
	// 如果鏈結串列是空的或只有一個節點，不需要反轉
	if isEmpty := l.IsEmpty(); isEmpty || l.head.next == nil {
		return
	}

	current := l.head
	var temp *Node[T]

	// 遍歷所有節點，反轉每個節點的 prev 和 next
	for current != nil {
		// 交換 prev 和 next
		temp = current.prev
		current.prev = current.next
		current.next = temp

		// 移動到下一個節點（原來的 prev，現在變成 next）
		current = current.prev
	}

	// 反轉後，更新 head 和 tail
	l.tail, l.head = l.head, l.tail
}

func (l *LinkedList[T]) IsEmpty() bool{
	if l.head == nil || l.tail == nil || l.length == 0 {
		return true
	}

	return false
}

func (l *LinkedList[T]) Merge(other *LinkedList[T]) {
	if otherIsEmpty := other.IsEmpty(); otherIsEmpty {
		return
	}

	if thisIsEmpty := l.IsEmpty(); thisIsEmpty {
		l.head = other.head
		l.tail = other.tail
		l.length = other.length
		return
	}

	l.tail.next = other.head
	other.head.prev = l.tail

	l.tail = other.tail
	l.length += other.length
}

func (l LinkedList[T]) PrintListData() {
	node := l.head

	if l.length == 0 || node == nil {
		fmt.Println("Empty Linked list...")
		return
	}

	for node != nil {
		if node.next == nil {
			fmt.Printf("%v\n", node.Data)
		} else {
			fmt.Printf("%v -> \t", node.Data)
		}

		node = node.next
	}
}
