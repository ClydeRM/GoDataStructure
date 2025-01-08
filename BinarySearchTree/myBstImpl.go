package BinarySearchTree

func (bst *BST[T]) Insert(data T) *Node[T] {
	newNode := NewNode(data)

	if bst.root == nil {
		bst.root = newNode
		return newNode
	}

	current := bst.root
	for {
		if data < current.Data {
			// 插入左子樹
			if current.Left == nil {
				current.Left = newNode
				break
			}
			current = current.Left
		} else {
			// 插入右子樹
			if current.Right == nil {
				current.Right = newNode
				break
			}
			current = current.Right
		}
	}

	return newNode
}