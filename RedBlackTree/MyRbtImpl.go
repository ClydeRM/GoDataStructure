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

// Insert
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
