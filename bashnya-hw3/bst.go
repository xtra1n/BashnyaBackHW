package bst

type BinarySearchTree[V Ordered] struct {
	root *Node[V]
	size uint
}

func New[V Ordered]() *BinarySearchTree[V] {
	return &BinarySearchTree[V]{
		root: nil,
		size: 0,
	}
}

func (bst *BinarySearchTree[V]) Insert(value V) {
	bst.root = bst.insertNode(bst.root, value)
}

func (bst *BinarySearchTree[V]) insertNode(node *Node[V], value V) *Node[V] {
	if node == nil {
		bst.size++
		return NewNode(value)
	}

	if value < node.Value {
		node.Left = bst.insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.insertNode(node.Right, value)
	} else {
		node.Value = value
	}

	return node
}

func (bst *BinarySearchTree[V]) Search(value V) bool {
	return bst.searchNode(bst.root, value)
}

func (bst *BinarySearchTree[V]) searchNode(node *Node[V], value V) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	} else if value < node.Value {
		return bst.searchNode(node.Left, value)
	} else {
		return bst.searchNode(node.Right, value)
	}
}

func (bst *BinarySearchTree[V]) Remove(value V) bool {
	var isSearched bool
	bst.root, isSearched = bst.removeNode(bst.root, value)

	if isSearched {
		bst.size--
	}

	return isSearched
}

func (bst BinarySearchTree[V]) removeNode(node *Node[V], value V) (*Node[V], bool) {
	if node == nil {
		return nil, false
	}

	var isSearched bool

	if node.Value > value {
		node.Left, isSearched = bst.removeNode(node.Left, value)
	} else if node.Value < value {
		node.Right, isSearched = bst.removeNode(node.Right, value)
	} else {
		isSearched = true

		if node.Left == nil && node.Right == nil {
			return nil, isSearched
		} else if node.Left == nil {
			return node.Right, isSearched
		} else if node.Right == nil {
			return node.Left, isSearched
		}

		minRight := bst.searchMin(node.Right)
		node.Value = minRight.Value
		node.Right, _ = bst.removeNode(node.Right, minRight.Value)
	}

	return node, isSearched
}

func (bst BinarySearchTree[V]) Min() (V, bool) {
	if bst.root == nil {
		var zero V
		return zero, false
	}

	node := bst.searchMin(bst.root)

	return node.Value, true
}

func (bst BinarySearchTree[V]) searchMin(node *Node[V]) *Node[V] {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (bst *BinarySearchTree[V]) Depth() int {
	return bst.depth(bst.root)
}

func (bst *BinarySearchTree[V]) depth(node *Node[V]) int {
	if node == nil {
		return -1
	}

	leftHeight := bst.depth(node.Left)
	rightHeight := bst.depth(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
