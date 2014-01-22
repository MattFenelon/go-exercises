package binarysearchtree

// A BinarySearchTree represents a collection of key-value-pairs
// stored in a tree structure.
// Child nodes to the left have keys less than the parent node and
// child nodes to the right have keys greater than the parent node.
// Getting and setting a key takes O(log n) time, however worst-case
// performance can be O(n) if the tree is severely unbalanced.
type BinarySearchTree struct {
	rootNode *node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Get traverses the BST to find the value for the
// specified key. Nil is returned if the key does not
// exist.
func (bst *BinarySearchTree) Get(key string) interface{} {
	return bst.rootNode.Get(key)
}

// Set stores the value against the key specified.
// If the key is already in use the value is overwritten with
// the value specified here.
func (bst *BinarySearchTree) Set(key string, value interface{}) {
	if bst.rootNode == nil {
		bst.rootNode = newNode(key, value)
		return
	}

	bst.rootNode.Set(key, value)
}

// node represents a node in the BST.
// A node can have between 0 and 2 child nodes, where the key of
// the left node is less than the parent key, and the key of the
// right node is greater than the parent key.
type node struct {
	key       string
	value     interface{}
	leftNode  *node
	rightNode *node
}

func newNode(key string, value interface{}) *node {
	return &node{key: key, value: value}
}

func (n *node) Get(key string) interface{} {
	if n == nil {
		return nil
	}

	if n.key == key {
		return n.value
	}

	if key > n.key {
		return n.rightNode.Get(key)
	}

	return n.leftNode.Get(key)
}

func (n *node) Set(key string, value interface{}) {
	if n.key == key {
		n.key, n.value = key, value
		return
	}

	traverseRightSide := key > n.key
	if traverseRightSide {
		if n.rightNode == nil {
			n.rightNode = newNode(key, value)
			return
		}

		n.rightNode.Set(key, value)
		return
	}

	if n.leftNode == nil {
		n.leftNode = newNode(key, value)
		return
	}

	n.leftNode.Set(key, value)
}
