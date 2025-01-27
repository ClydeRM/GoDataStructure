package RedBlackTree

func (rbt *RBT[T]) IsEmpty() bool {
    return rbt.root == nil
}

func (rbt *RBT[T]) Size() int {
    return rbt.size
}

func (rbt *RBT[T]) Min() (T, bool) {
    if rbt.IsEmpty() {
        var zero T
        return zero, false
    }
    node := rbt.root
    for node.left != nil {
        node = node.left
    }
    return node.Data, true
}

func (rbt *RBT[T]) Max() (T, bool) {
    if rbt.IsEmpty() {
        var zero T
        return zero, false
    }
    node := rbt.root
    for node.right != nil {
        node = node.right
    }
    return node.Data, true
}

func (rbt *RBT[T]) InOrderTraversal(root *Node[T]) []T {
    // 減少Slice 的複製跟append效能消耗
    return rbt.inOrderRecursively(root, []T{})
}

func (rbt *RBT[T]) inOrderRecursively(node *Node[T], result []T) []T {
    if node == nil {
		return result
	}

    // LVR
	result = rbt.inOrderRecursively(node.left, result)

    result = append(result, node.Data)

    result = rbt.inOrderRecursively(node.right, result)

	return result
}

func (rbt *RBT[T]) Insert(data T) {
	// step 1. BST insert node
	newNode := NewNode(data, RED) // default Color is red.
	newNode.parent = nil
	newNode.left = rbt.nilNode
	newNode.right = rbt.nilNode

	// step 2. find insert place
	parent := rbt.nilNode
	current := rbt.root

	for current != rbt.nilNode {
		parent = current
		if data < current.Data {
			current = current.left
		} else {
			current = current.right
		}
	}

	// step 3. insert newNode
	if parent == rbt.nilNode {
		rbt.root = newNode
	} else if data < parent.Data {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	newNode.parent = parent
	rbt.size++

	// step 4. fix up RBT
	rbt.fixInsert(newNode)
}

// fixInsert
func (rbt *RBT[T]) fixInsert(current *Node[T]) {
	// current.parent is RED: violate RBT rule "RED by RED"
	// Iterative maintain RBT feature until parent change to BLACK.
	for current.parent != nil && current.parent.Color == RED {
		// step 1. find grandparent to find uncle
		grandparent := current.parent.parent

		if current.parent == grandparent.left {
			// step 2. find uncle
			uncle := grandparent.right

			// case 1. uncle is RED
			if uncle.Color == RED {
				// reset parent and grandparent color
				// parent -> black, uncle -> black, grandparent -> red
				current.parent.Color = BLACK
				uncle.Color = BLACK
				grandparent.Color = RED
				current = grandparent // iterative grandparent
			} else {
				// case 2& 3. uncle is black
				// case 2. current is right child
				if current == current.parent.right {
					// left rotate parent
					current = current.parent
					rbt.leftRotate(current)
				}
				// case 3. current is left child
				// parent -> black, grandparent -> red + right rotate grand parent
				current.parent.Color = BLACK
				grandparent.Color = RED
				rbt.rightRotate(grandparent)
			}
		} else {
			// symmetry to above.
			uncle := grandparent.left

			// case 1. uncle is RED
			if uncle.Color == RED {
				current.parent.Color = BLACK
				uncle.Color = BLACK
				grandparent.Color = RED
				current = grandparent
			} else {
				// case 2& 3. uncle is black
				// case 2. current is right child
				if current == current.parent.left {
					// right rotate parent
					current = current.parent
					rbt.rightRotate(current)
				}
				// case 3. current is left child
				// parent -> black, grandparent -> red + left rotate grand parent
				current.parent.Color = BLACK
				grandparent.Color = RED
				rbt.leftRotate(grandparent)
			}
		}
	}

	// root always black
	rbt.root.Color = BLACK
}

// leftRotate
func (rbt *RBT[T]) leftRotate(target *Node[T]) {
	// step 1. hoist rightChild
	rightChild := target.right
	target.right = rightChild.left

	if rightChild.left != rbt.nilNode {
		rightChild.left.parent = target
	}

	// step 2. rightChild parent reset
	rightChild.parent = target.parent
	if target.parent == nil {
		rbt.root = rightChild
	} else if target == target.parent.left {
		target.parent.left = rightChild
	} else {
		target.parent.right = rightChild
	}

	// step 3. rightChild left reset
	rightChild.left = target
	target.parent = rightChild
}

// rightRotate
func (rbt *RBT[T]) rightRotate(target *Node[T]) {
	// step 1. hosit leftChild
	leftChild := target.left
	target.left = leftChild.right
	if leftChild.right != rbt.nilNode {
		leftChild.right.parent = target
	}

	// step 2. leftChild parent reset
	leftChild.parent = target.parent
	if target.parent == nil {
		rbt.root = leftChild
	} else if target == target.parent.right {
		target.parent.right = leftChild
	} else {
		target.parent.left = leftChild
	}

	// step 3. leftChild right reset
	leftChild.right = target
	target.parent = leftChild
}

func (rbt *RBT[T]) Delete(data T) {
	// 找到節點
	node := rbt.Search(rbt.root, data)
	if node == rbt.nilNode {
		return // 節點不存在
	}

	// 執行刪除
	rbt.deleteNode(node)
}

// 找到指定數據的節點
func (rbt *RBT[T]) Search(node *Node[T], data T) *Node[T] {
	for node != rbt.nilNode {
		if data == node.Data {
			return node
		} else if data < node.Data {
			node = node.left
		} else {
			node = node.right
		}
	}
	return rbt.nilNode
}

// 刪除節點邏輯
func (rbt *RBT[T]) deleteNode(target *Node[T]) {
	var child, toFix *Node[T]
	originalColor := target.Color

	if target.left == rbt.nilNode {
		// case 1: target no left child
		child = target.right
		rbt.transplant(target, target.right)
		toFix = child
	} else if target.right == rbt.nilNode {
		// case 2: target no right child
		child = target.left
		rbt.transplant(target, target.left)
		toFix = child
	} else {
		// case 3: target have both children
		successor := rbt.minimum(target.right) // 找到後繼節點
		originalColor = successor.Color
		child = successor.right
		toFix = child

		if successor.parent == target {
			child.parent = successor
		} else {
			rbt.transplant(successor, successor.right)
			successor.right = target.right
			successor.right.parent = successor
		}
		rbt.transplant(target, successor)
		successor.left = target.left
		successor.left.parent = successor
		successor.Color = target.Color
	}

	// fix rbt property, when target is black
	if originalColor == BLACK {
		rbt.fixDelete(toFix)
	}
}

// 找到子樹中最小值節點
func (rbt *RBT[T]) minimum(node *Node[T]) *Node[T] {
	for node.left != rbt.nilNode {
		node = node.left
	}
	return node
}

// 替換節點
func (rbt *RBT[T]) transplant(u, v *Node[T]) {
	if u.parent == nil {
		rbt.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (rbt *RBT[T]) fixDelete(node *Node[T]) {
	for node != rbt.root && node.Color == BLACK {
		if node == node.parent.left {
			sibling := node.parent.right

			// case 1: sibling is red
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.parent.Color = RED
				rbt.leftRotate(node.parent)
				sibling = node.parent.right
			}

			// case 2: sibling is black & both children are black
			if sibling.left.Color == BLACK && sibling.right.Color == BLACK {
				sibling.Color = RED
				node = node.parent
			} else {
				// case 3: right is black, left is red
				if sibling.right.Color == BLACK {
					sibling.left.Color = BLACK
					sibling.Color = RED
					rbt.rightRotate(sibling)
					sibling = node.parent.right
				}
				// case 4: both are red
				sibling.Color = node.parent.Color
				node.parent.Color = BLACK
				sibling.right.Color = BLACK
				rbt.leftRotate(node.parent)
				node = rbt.root
			}
		} else {
			// symmetry above
			sibling := node.parent.left

			// case 1: sibling is red
			if sibling.Color == RED {
				sibling.Color = BLACK
				node.parent.Color = RED
				rbt.rightRotate(node.parent)
				sibling = node.parent.left
			}

			// case 2: sibling is black and both children are black
			if sibling.right.Color == BLACK && sibling.left.Color == BLACK {
				sibling.Color = RED
				node = node.parent
			} else {
				// case 3: left is black, right is red
				if sibling.left.Color == BLACK {
					sibling.right.Color = BLACK
					sibling.Color = RED
					rbt.leftRotate(sibling)
					sibling = node.parent.left
				}
				// case 4: both are red
				sibling.Color = node.parent.Color
				node.parent.Color = BLACK
				sibling.left.Color = BLACK
				rbt.rightRotate(node.parent)
				node = rbt.root
			}
		}
	}
	node.Color = BLACK
}


func (rbt *RBT[T]) DeleteRecursively(data T) *Node[T] {
	if rbt.IsEmpty() {
		return rbt.root
	}

	rbt.root = rbt.deleteRecursively(rbt.root, data)
	rbt.root.parent = rbt.nilNode
	rbt.root.Color = BLACK
	return rbt.root
}

func (rbt *RBT[T]) deleteRecursively(node *Node[T], data T) *Node[T]{
	if node == rbt.nilNode {
		return rbt.nilNode
	}

	if data < node.Data {
		node.left = rbt.deleteRecursively(node.left, data)
	} else if data > node.Data {
		node.right = rbt.deleteRecursively(node.right, data)
	} else {
		// find the target
		if node.left == rbt.nilNode && node.right == rbt.nilNode {
			// Case 1: target is leaf
			return rbt.nilNode
		} else if node.left == rbt.nilNode {
			// Case 2: target only have subright
			return node.right
		} else if node.right == rbt.nilNode {
			// Case 3: target only have subleft
			return node.left
		} else {
			// Case 4: target have both
			successor := rbt.minimum(node.right) // 找到後繼節點
			node.Data = successor.Data
			node.right = rbt.deleteRecursively(node.right, successor.Data)
		}
	}

	// fix rbt property
	return rbt.fixDeleteRecursively(node)
}

func (rbt *RBT[T]) fixDeleteRecursively(node *Node[T]) *Node[T] {
	if node.Color == RED {
		return node
	}

	// if the delete node is black, fix rbt property
	// node is left child
	if node == node.parent.left {
		sibling := node.parent.right

		// Case 1: sibling is red
		if sibling.Color == RED {
			sibling.Color = BLACK
			node.parent.Color = RED
			rbt.leftRotate(node.parent)
			sibling = node.parent.right
		}

		// Case 2: sibling is black, and both children are black
		if sibling.left.Color == BLACK && sibling.right.Color == BLACK {
			sibling.Color = RED
			return node.parent
		}

		// Case 3: right is black, left is red
		if sibling.right.Color == BLACK {
			sibling.left.Color = BLACK
			sibling.Color = RED
			rbt.rightRotate(sibling)
			sibling = node.parent.right
		}

		// Case 4: both are red
		sibling.Color = node.parent.Color
		node.parent.Color = BLACK
		sibling.right.Color = BLACK
		rbt.leftRotate(node.parent)
	} else {
		// symmetry above
		sibling := node.parent.left

		// Case 1: sibling is red
		if sibling.Color == RED {
			sibling.Color = BLACK
			node.parent.Color = RED
			rbt.rightRotate(node.parent)
			sibling = node.parent.left
		}

		// Case 2: sibling is black, and both children are black
		if sibling.left.Color == BLACK && sibling.right.Color == BLACK {
			sibling.Color = RED
			return node.parent
		}

		// Case 3: left is black, right is red
		if sibling.left.Color == BLACK {
			sibling.right.Color = BLACK
			sibling.Color = RED
			rbt.leftRotate(sibling)
			sibling = node.parent.left
		}

		// Case 4: both are red
		sibling.Color = node.parent.Color
		node.parent.Color = BLACK
		sibling.left.Color = BLACK
		rbt.rightRotate(node.parent)
	}
	return node
}
