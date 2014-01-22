package binarysearchtree

type BinarySearchTree struct {
	rootNode *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	key       string
	value     interface{}
	leftNode  *BinarySearchTreeNode
	rightNode *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		rootNode: nil,
	}
}

func (bst *BinarySearchTree) Get(key string) interface{} {
	return findValueForKey(bst.rootNode, key)
}

func (bst *BinarySearchTree) Set(key string, value interface{}) {
	if bst.rootNode == nil {
		bst.rootNode = &BinarySearchTreeNode{key: key, value: value}
		return
	}

	n := findNodeForKey(bst.rootNode, key)
	n.value = value
}

func findValueForKey(current *BinarySearchTreeNode, key string) interface{} {
	if current == nil {
		return nil
	}

	if current.key == key {
		return current.value
	}

	if key > current.key {
		return findValueForKey(current.rightNode, key)
	}

	return findValueForKey(current.leftNode, key)
}

func findNodeForKey(current *BinarySearchTreeNode, key string) *BinarySearchTreeNode {
	if key == current.key {
		return current
	}

	traverseRightSide := key > current.key
	if traverseRightSide {
		if current.rightNode == nil {
			current.rightNode = &BinarySearchTreeNode{key: key}
			return current.rightNode
		}

		return findNodeForKey(current.rightNode, key)
	}

	if current.leftNode == nil {
		current.leftNode = &BinarySearchTreeNode{key: key}
		return current.leftNode
	}

	return findNodeForKey(current.leftNode, key)
}
