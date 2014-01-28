package redblacktree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_when_adding_a_node_to_an_empty_tree(t *testing.T) {
	actualRbt := NewRedBlackTree()

	actualRbt.Set("a", "a")

	expectedRbt := &RedBlackTree{
		root: &RedBlackTreeNode{key: "a", value: "a", colour: black}}

	t.Log("It should set the node as the root and colour it black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_node_is_added_to_root(t *testing.T) {
	t.Log("And the key is greater than the root")
	actualRbt := NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")

	expectedRbt := &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "a",
			value:  "a",
			colour: black,
			right:  &RedBlackTreeNode{key: "c", value: "c", colour: red},
		},
	}

	t.Log("It should add the node as the right child red node of the black root")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than the root")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("c", "c")
	actualRbt.Set("a", "a")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left:   &RedBlackTreeNode{key: "a", value: "a", colour: red},
		},
	}

	t.Log("It should add the node as the left child red node of the black root")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_node_is_added_to_root_with_one_child_node(t *testing.T) {
	t.Log("And the sibling is less than the root and the new key is greater than the root")
	actualRbt := NewRedBlackTree()

	actualRbt.Set("c", "c")
	actualRbt.Set("a", "a")
	actualRbt.Set("e", "e")

	expectedRbt := &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left:   &RedBlackTreeNode{key: "a", value: "a", colour: red},
			right:  &RedBlackTreeNode{key: "e", value: "e", colour: red},
		},
	}

	t.Log("It should add the new node as the right child red node of the black root")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the sibling is greater than the root and the new key is less than the root")
	actualRbt = NewRedBlackTree()

	actualRbt.Set("c", "c")
	actualRbt.Set("e", "e")
	actualRbt.Set("a", "a")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left:   &RedBlackTreeNode{key: "a", value: "a", colour: red},
			right:  &RedBlackTreeNode{key: "e", value: "e", colour: red},
		},
	}

	t.Log("It should add the new node as the left child red node of the black root")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_node_is_added_to_a_red_parent_and_its_uncle_is_red_and_its_grandparent_is_root(t *testing.T) {
	t.Log("And the new key is greater than the root and its parent")
	actualRbt := NewRedBlackTree()

	actualRbt.Set("key5", "value5")
	actualRbt.Set("key2", "value2")
	actualRbt.Set("key8", "value8")
	actualRbt.Set("key9", "value9")

	expectedRbt := &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "key5",
			value:  "value5",
			colour: black,
			left:   &RedBlackTreeNode{key: "key2", value: "value2", colour: black},
			right: &RedBlackTreeNode{
				key:    "key8",
				value:  "value8",
				colour: black,
				right:  &RedBlackTreeNode{key: "key9", value: "value9", colour: red},
			},
		},
	}

	t.Log("It should add the node as the right child of its parent, flip the colours of the parent and uncle to black, and leave root as black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is greater than the root and less than its parent")
	actualRbt = NewRedBlackTree()

	actualRbt.Set("key5", "value5")
	actualRbt.Set("key2", "value2")
	actualRbt.Set("key8", "value8")
	actualRbt.Set("key7", "value7")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "key5",
			value:  "value5",
			colour: black,
			left:   &RedBlackTreeNode{key: "key2", value: "value2", colour: black},
			right: &RedBlackTreeNode{
				key:    "key8",
				value:  "value8",
				colour: black,
				left:   &RedBlackTreeNode{key: "key7", value: "value7", colour: red},
			},
		},
	}

	t.Log("It should add the node as the left child of its parent, flip the colours of the parent and uncle to black, and leave root as black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is less than the root and greater than its parent")
	actualRbt = NewRedBlackTree()

	actualRbt.Set("key5", "value5")
	actualRbt.Set("key2", "value2")
	actualRbt.Set("key8", "value8")
	actualRbt.Set("key3", "value3")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "key5",
			value:  "value5",
			colour: black,
			left: &RedBlackTreeNode{
				key:    "key2",
				value:  "value2",
				colour: black,
				right:  &RedBlackTreeNode{key: "key3", value: "value3", colour: red}},
			right: &RedBlackTreeNode{key: "key8", value: "value8", colour: black},
		},
	}

	t.Log("It should add the node as the right child of its parent, flip the colours of the parent and uncle to black, and leave root as black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is less than the root and less than its parent")
	actualRbt = NewRedBlackTree()

	actualRbt.Set("key5", "value5")
	actualRbt.Set("key2", "value2")
	actualRbt.Set("key8", "value8")
	actualRbt.Set("key1", "value1")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "key5",
			value:  "value5",
			colour: black,
			left: &RedBlackTreeNode{
				key:    "key2",
				value:  "value2",
				colour: black,
				left:   &RedBlackTreeNode{key: "key1", value: "value1", colour: red}},
			right: &RedBlackTreeNode{key: "key8", value: "value8", colour: black},
		},
	}

	t.Log("It should add the node as the left child of its parent, flip the colours of the parent and uncle to black, and leave root as black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_node_is_added_to_a_red_parent_and_its_uncle_is_red_and_its_grandparent_is_not_root(t *testing.T) {
	newTestTree := func() *RedBlackTree {
		return &RedBlackTree{
			root: &RedBlackTreeNode{
				key:    "j",
				value:  "j",
				colour: black,
				left: &RedBlackTreeNode{
					key: "e", value: "e", colour: black,
					left:  &RedBlackTreeNode{key: "b", value: "b", colour: red},
					right: &RedBlackTreeNode{key: "g", value: "g", colour: red},
				},
				right: &RedBlackTreeNode{
					key: "o", value: "o", colour: black,
					left:  &RedBlackTreeNode{key: "l", value: "l", colour: red},
					right: &RedBlackTreeNode{key: "v", value: "v", colour: red},
				},
			},
		}
	}

	createActualTree := func() *RedBlackTree {
		actualRbt := NewRedBlackTree()

		actualRbt.Set("j", "j")
		actualRbt.Set("e", "e")
		actualRbt.Set("o", "o")
		actualRbt.Set("b", "b")
		actualRbt.Set("g", "g")
		actualRbt.Set("l", "l")
		actualRbt.Set("v", "v")

		return actualRbt
	}

	t.Log("And the new key is less than the root, its grandparent and its parent")
	actualRbt := createActualTree()
	actualRbt.Set("a", "a")

	expectedRbt := newTestTree()
	expectedRbt.root.left.colour = red
	expectedRbt.root.left.left.colour = black
	expectedRbt.root.left.right.colour = black
	expectedRbt.root.left.left.left = &RedBlackTreeNode{key: "a", value: "a", colour: red}

	t.Log("It should add the node as the left child of the leftmost node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is less than the root and its grandparent, but greater than its parent")
	actualRbt = createActualTree()
	actualRbt.Set("c", "c")

	expectedRbt = newTestTree()
	expectedRbt.root.left.colour = red
	expectedRbt.root.left.left.colour = black
	expectedRbt.root.left.right.colour = black
	expectedRbt.root.left.left.right = &RedBlackTreeNode{key: "c", value: "c", colour: red}

	t.Log("It should add the node as the right child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is less than the root, greater than its grandparent and less than its parent")
	actualRbt = createActualTree()
	actualRbt.Set("f", "f")

	expectedRbt = newTestTree()
	expectedRbt.root.left.colour = red
	expectedRbt.root.left.left.colour = black
	expectedRbt.root.left.right.colour = black
	expectedRbt.root.left.right.left = &RedBlackTreeNode{key: "f", value: "f", colour: red}

	t.Log("It should add the node as the left child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is less than the root, greater than its grandparent and its parent")
	actualRbt = createActualTree()
	actualRbt.Set("h", "h")

	expectedRbt = newTestTree()
	expectedRbt.root.left.colour = red
	expectedRbt.root.left.left.colour = black
	expectedRbt.root.left.right.colour = black
	expectedRbt.root.left.right.right = &RedBlackTreeNode{key: "h", value: "h", colour: red}

	t.Log("It should add the node as the right child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is greater than the root but less than its grandparent and parent")
	actualRbt = createActualTree()
	actualRbt.Set("k", "k")

	expectedRbt = newTestTree()
	expectedRbt.root.right.colour = red
	expectedRbt.root.right.left.colour = black
	expectedRbt.root.right.right.colour = black
	expectedRbt.root.right.left.left = &RedBlackTreeNode{key: "k", value: "k", colour: red}

	t.Log("It should add the node as the left child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is greater than the root, less than its grandparent and greater than its parent")
	actualRbt = createActualTree()
	actualRbt.Set("m", "m")

	expectedRbt = newTestTree()
	expectedRbt.root.right.colour = red
	expectedRbt.root.right.left.colour = black
	expectedRbt.root.right.right.colour = black
	expectedRbt.root.right.left.right = &RedBlackTreeNode{key: "m", value: "m", colour: red}

	t.Log("It should add the node as the right child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is greater than the root and its grandparent, but less than its parent")
	actualRbt = createActualTree()
	actualRbt.Set("p", "p")

	expectedRbt = newTestTree()
	expectedRbt.root.right.colour = red
	expectedRbt.root.right.left.colour = black
	expectedRbt.root.right.right.colour = black
	expectedRbt.root.right.right.left = &RedBlackTreeNode{key: "p", value: "p", colour: red}

	t.Log("It should add the node as the left child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the new key is greater than the root, its grandparent and its parent")
	actualRbt = createActualTree()
	actualRbt.Set("x", "x")

	expectedRbt = newTestTree()
	expectedRbt.root.right.colour = red
	expectedRbt.root.right.left.colour = black
	expectedRbt.root.right.right.colour = black
	expectedRbt.root.right.right.right = &RedBlackTreeNode{key: "x", value: "x", colour: red}

	t.Log("It should add the node as the right child of its parent node and flip the colours of its parent and uncle to black and its grandparent to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_new_node_is_added_to_a_red_parent_with_no_uncle_and_its_grandparent_is_root(t *testing.T) {
	t.Log("And the key is greater than root and its parent")
	actualRbt := NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")
	actualRbt.Set("e", "e")

	expectedRbt := &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: red,
			},
			right: &RedBlackTreeNode{
				key: "e", value: "e", colour: red,
			},
		},
	}

	t.Log("It should left rotate on the root and flip the colour of the new root to black and the old root to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is greater than root and less than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("e", "e")
	actualRbt.Set("c", "c")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: red,
			},
			right: &RedBlackTreeNode{
				key: "e", value: "e", colour: red,
			},
		},
	}

	t.Log("It should right rotate on its parent, then left rotate on its grandparent (root), and flip the colour of the new root to black and the old root to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than root and its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("e", "e")
	actualRbt.Set("c", "c")
	actualRbt.Set("a", "a")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: red,
			},
			right: &RedBlackTreeNode{
				key: "e", value: "e", colour: red,
			},
		},
	}

	t.Log("It should right rotate on the root and flip the colour of the new root to black and the old root to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than root and greater than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("e", "e")
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: red,
			},
			right: &RedBlackTreeNode{
				key: "e", value: "e", colour: red,
			},
		},
	}

	t.Log("It should left rotate on its parent, then right rotate on its grandparent (root) and flip the colour of the new root to black and the old root to red")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_new_node_is_added_to_a_red_parent_with_no_uncle_and_its_grandparent_is_not_root(t *testing.T) {
	t.Log("And the key is greater than root, its grandparent, and its parent")
	actualRbt := NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")
	actualRbt.Set("e", "e")
	actualRbt.Set("h", "h")
	actualRbt.Set("l", "l")

	expectedRbt := &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: black,
			},
			right: &RedBlackTreeNode{
				key: "h", value: "h", colour: black,
				left:  &RedBlackTreeNode{key: "e", value: "e", colour: red},
				right: &RedBlackTreeNode{key: "l", value: "l", colour: red},
			},
		},
	}

	t.Log("It should left rotate on its grandparent and flip the colour of its grandparent to red and its parent to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is greater than root and its grandparent, and less than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")
	actualRbt.Set("e", "e")
	actualRbt.Set("l", "l")
	actualRbt.Set("h", "h")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: black,
			},
			right: &RedBlackTreeNode{
				key: "h", value: "h", colour: black,
				left:  &RedBlackTreeNode{key: "e", value: "e", colour: red},
				right: &RedBlackTreeNode{key: "l", value: "l", colour: red},
			},
		},
	}

	t.Log("It should right rotate on its parent, then left rotate on its grandparent, and flip the colour of its grandparent to red and its colour to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is greater than root, less than its grandparent, and less than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")
	actualRbt.Set("l", "l")
	actualRbt.Set("h", "h")
	actualRbt.Set("e", "e")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: black,
			},
			right: &RedBlackTreeNode{
				key: "h", value: "h", colour: black,
				left:  &RedBlackTreeNode{key: "e", value: "e", colour: red},
				right: &RedBlackTreeNode{key: "l", value: "l", colour: red},
			},
		},
	}

	t.Log("It should right rotate on its grandparent, and flip the colour of its grandparent to red and its parent to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is greater than root, less than its grandparent, and greater than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")
	actualRbt.Set("l", "l")
	actualRbt.Set("e", "e")
	actualRbt.Set("h", "h")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "c",
			value:  "c",
			colour: black,
			left: &RedBlackTreeNode{
				key: "a", value: "a", colour: black,
			},
			right: &RedBlackTreeNode{
				key: "h", value: "h", colour: black,
				left:  &RedBlackTreeNode{key: "e", value: "e", colour: red},
				right: &RedBlackTreeNode{key: "l", value: "l", colour: red},
			},
		},
	}

	t.Log("It should left rotate on its parent, then right rotate on its grandparent, and flip its grandparent's colour to red, and its colour to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than root, its grandparent, and its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("l", "l")
	actualRbt.Set("h", "h")
	actualRbt.Set("e", "e")
	actualRbt.Set("c", "c")
	actualRbt.Set("a", "a")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "h",
			value:  "h",
			colour: black,
			left: &RedBlackTreeNode{
				key: "c", value: "c", colour: black,
				left:  &RedBlackTreeNode{key: "a", value: "a", colour: red},
				right: &RedBlackTreeNode{key: "e", value: "e", colour: red},
			},
			right: &RedBlackTreeNode{
				key: "l", value: "l", colour: black,
			},
		},
	}

	t.Log("It should right rotate on its grandparent, and flip the colours of its grandparent's to red and its parents to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than root, its grandparent, and greater than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("l", "l")
	actualRbt.Set("h", "h")
	actualRbt.Set("e", "e")
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "h",
			value:  "h",
			colour: black,
			left: &RedBlackTreeNode{
				key: "c", value: "c", colour: black,
				left:  &RedBlackTreeNode{key: "a", value: "a", colour: red},
				right: &RedBlackTreeNode{key: "e", value: "e", colour: red},
			},
			right: &RedBlackTreeNode{
				key: "l", value: "l", colour: black,
			},
		},
	}

	t.Log("It should left rotate on its parent, then right rotate on its grandparent, and flip its grandparent's colour to red, and its colour to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than root, greater than its grandparent, and less than its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("l", "l")
	actualRbt.Set("h", "h")
	actualRbt.Set("a", "a")
	actualRbt.Set("e", "e")
	actualRbt.Set("c", "c")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "h",
			value:  "h",
			colour: black,
			left: &RedBlackTreeNode{
				key: "c", value: "c", colour: black,
				left:  &RedBlackTreeNode{key: "a", value: "a", colour: red},
				right: &RedBlackTreeNode{key: "e", value: "e", colour: red},
			},
			right: &RedBlackTreeNode{
				key: "l", value: "l", colour: black,
			},
		},
	}

	t.Log("It should right rotate on its parent, and then left rotate on its grandparent, and flip its grandparent's colour to red and its colour to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}

	t.Log("And the key is less than root, greater than its grandparent and its parent")
	actualRbt = NewRedBlackTree()
	actualRbt.Set("l", "l")
	actualRbt.Set("h", "h")
	actualRbt.Set("a", "a")
	actualRbt.Set("c", "c")
	actualRbt.Set("e", "e")

	expectedRbt = &RedBlackTree{
		root: &RedBlackTreeNode{
			key:    "h",
			value:  "h",
			colour: black,
			left: &RedBlackTreeNode{
				key: "c", value: "c", colour: black,
				left:  &RedBlackTreeNode{key: "a", value: "a", colour: red},
				right: &RedBlackTreeNode{key: "e", value: "e", colour: red},
			},
			right: &RedBlackTreeNode{
				key: "l", value: "l", colour: black,
			},
		},
	}

	t.Log("It should left rotate on its grandparent, and flip the colour of its grandparent to red and its parent to black")
	if reflect.DeepEqual(expectedRbt, actualRbt) == false {
		t.Errorf("\tExpected %v but was %v", expectedRbt, actualRbt)
	}
}

func Test_when_a_node_is_added_to_a_red_parent_and_its_uncle_is_black_and_its_grandparent_is_root(t *testing.T) {
	t.Skip()
}

func Test_when_a_node_is_added_to_a_red_parent_and_its_uncle_is_black_and_its_grandparent_is_not_root(t *testing.T) {
	t.Skip()
}

type RedBlackTree struct {
	root *RedBlackTreeNode
}

type RedBlackTreeNode struct {
	key         string
	value       interface{}
	colour      NodeColour
	left, right *RedBlackTreeNode
}

type NodeColour bool

const (
	black NodeColour = false
	red   NodeColour = true
)

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (rbt *RedBlackTree) Set(key string, value interface{}) {
	if rbt.root == nil {
		rbt.root = &RedBlackTreeNode{key: key, value: value}
		return
	}

	n := rbt.root
	insertNewNode(n, nil, nil, key, value)
}

func insertNewNode(n, parent, grandparent *RedBlackTreeNode, key string, value interface{}) {
	if key > n.key {
		if n.right == nil {
			n.right = &RedBlackTreeNode{key: key, value: value, colour: red}
			rebalanceNode(n.right, n, parent)
			return
		}

		insertNewNode(n.right, n, parent, key, value)
		rebalanceNode(n, parent, grandparent)
		return
	}

	if n.left == nil {
		n.left = &RedBlackTreeNode{key: key, value: value, colour: red}
		rebalanceNode(n.left, n, parent)
		return
	}

	insertNewNode(n.left, n, parent, key, value)
	rebalanceNode(n, parent, grandparent)
}

func rebalanceNode(n, parent, grandparent *RedBlackTreeNode) {
	if parent == nil {
		n.colour = black
		return
	}

	if grandparent == nil {
		return
	}

	var uncle *RedBlackTreeNode
	if parent == grandparent.right {
		uncle = grandparent.left
	} else {
		uncle = grandparent.right
	}

	if uncle == nil && parent.colour == red && n == parent.right && parent == grandparent.right {
		parent.colour = black
		grandparent.colour = red

		// left rotate on grandparent
		grandparent.right = nil
		saved := *grandparent
		*grandparent = *parent
		grandparent.left = &saved
	}

	if uncle == nil && parent.colour == red && n == parent.left && parent == grandparent.left {
		parent.colour = black
		grandparent.colour = red

		// right rotate on grandparent
		grandparent.left = nil
		saved := *grandparent
		*grandparent = *parent
		grandparent.right = &saved
	}

	if uncle == nil && parent.colour == red && n == parent.right && parent == grandparent.left {
		// left rotate on parent
		parent.right = nil
		n.left = parent
		grandparent.left = n
		rebalanceNode(n.left, n, grandparent) // could just be rotate-right
	}

	if uncle == nil && parent.colour == red && n == parent.left && parent == grandparent.right {
		// right rotate on parent
		parent.left = nil
		n.right = parent
		grandparent.right = n
		rebalanceNode(n.right, n, grandparent) // could just be rotate-left
	}

	if uncle != nil && uncle.colour == red && parent.colour == red {
		parent.colour = black
		uncle.colour = black
		grandparent.colour = red
	}
}

func (rbt *RedBlackTree) String() string {
	return fmt.Sprintf("%v", rbt.root)
}

func (n *RedBlackTreeNode) String() string {
	return fmt.Sprintf("[%v <- %#v:%#v (%v) -> %v]", n.left, n.key, n.value, n.colour, n.right)
}

func (n NodeColour) String() string {
	if n == red {
		return "red"
	}

	return "black"
}
