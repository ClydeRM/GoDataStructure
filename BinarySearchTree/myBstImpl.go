package BinarySearchTree

func (bst *BST[T]) Root() *Node[T] {
	return bst.root
}
func (bst *BST[T]) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BST[T]) Insert(data T) *Node[T] {
	newNode := NewNode(data)

	if bst.IsEmpty() {
		bst.root = newNode
		return newNode
	}

	current := bst.Root()
	for {
		if data < current.Data {
			// insert left subtree
			if current.left == nil {
				current.left = newNode
				newNode.parent = current
				break
			}
			current = current.left
		} else {
			// insert right subtree
			if current.right == nil {
				current.right = newNode
				newNode.parent = current
				break
			}
			current = current.right
		}
	}

	return newNode
}

func (bst *BST[T]) Search(data T) *Node[T] {
	if bst.IsEmpty() {
		return nil
	}

	current := bst.Root()
	for current != nil {
		if data == current.Data {
			return current
		} else if data < current.Data {
			current = current.left
		} else {
			current = current.right
		}
	}

	return nil
}

func (bst *BST[T]) PreOrderTraversal(root *Node[T]) []T {
	// VLR
	if root == nil {
		return []T{}
	}

	current := root
	var result []T
	result = append(result, current.Data)
	result = append(result, bst.PreOrderTraversal(current.left)...)
	result = append(result, bst.PreOrderTraversal(current.right)...)
	return result
}

func (bst *BST[T]) InOrderTraversal(root *Node[T]) []T {
	// LVR
	if root == nil {
		return []T{}
	}

	current := root
	var result []T
	result = append(result, bst.InOrderTraversal(current.left)...)
	result = append(result, current.Data)
	result = append(result, bst.InOrderTraversal(current.right)...)
	return result
}

func (bst *BST[T]) PostOrderTraversal(root *Node[T]) []T {
	// LRV
	if root == nil {
		return []T{}
	}

	current := root
	var result []T
	result = append(result, bst.PostOrderTraversal(current.left)...)
	result = append(result, bst.PostOrderTraversal(current.right)...)
	result = append(result, current.Data)
	return result
}

func (bst *BST[T]) Min(root *Node[T]) *Node[T] {
	current := root
	for current != nil && current.left != nil {
		current = current.left
	}
	return current
}

func (bst *BST[T]) Max(root *Node[T]) *Node[T] {
	current := root
	for current != nil && current.right != nil {
		current = current.right
	}
	return current
}

func (bst *BST[T]) InOrderSuccessor(node *Node[T]) *Node[T] {
	// LVR
	if node == nil {
		return nil
	}

	if node.right != nil {
		return bst.Min(node.right)
	}

	current := node
	successor := current.parent
	for successor != nil && current == successor.right {
		current = successor
		successor = successor.parent
	}
	return successor
}

func (bst *BST[T]) InOrderPredecessor(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	if node.left != nil {
		return bst.Max(node.left)
	}

	current := node
	predecessor := current.parent
	for predecessor != nil && current == predecessor.left {
		current = predecessor
		predecessor = predecessor.parent
	}
	return predecessor
}
