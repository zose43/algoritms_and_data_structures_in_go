package tree

import (
	"errors"
	"fmt"
)

type Rib interface {
	Key() int
}

type Node[T Rib] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func newNode[T Rib](data T) *Node[T] {
	return &Node[T]{data: data}
}

func (n *Node[T]) key() int {
	return n.data.Key()
}

func (n *Node[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

type BinaryTree[T Rib] struct {
	root *Node[T]
}

func NewBinaryTree[T Rib]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (t *BinaryTree[T]) Root() *Node[T] {
	return t.root
}

func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree[T]) Insert(Value T) {
	newEl := newNode[T](Value)
	if t.root == nil {
		t.root = newEl
	} else {
		current := t.root
		var parent *Node[T]
		for {
			parent = current
			if current.data.Key() > newEl.key() {
				current = current.left
				if current == nil {
					parent.left = newEl
					return
				}
			} else {
				current = current.right
				if current == nil {
					parent.right = newEl
					return
				}
			}
		}
	}
}

func (t *BinaryTree[T]) Find(key int) (T, bool) {
	current := t.root
	for current != nil {
		if current.data.Key() == key {
			return current.data, true
		}
		if current.data.Key() > key {
			current = current.left
		} else {
			current = current.right
		}
	}
	return *new(T), false
}

func (t *BinaryTree[T]) Minimum() (T, bool) {
	current := t.root
	var minimum *Node[T]
	for current != nil {
		minimum = current
		current = current.left
	}
	if minimum == nil {
		return *new(T), false
	}
	return minimum.data, true
}

func (t *BinaryTree[T]) Maximum() (T, bool) {
	current := t.root
	var maximum *Node[T]
	for current != nil {
		maximum = current
		current = current.right
	}
	if maximum == nil {
		return *new(T), false
	}
	return maximum.data, true
}

var ErrNotFoundElementByKey = errors.New("cannot find element by key")

func (t *BinaryTree[T]) Remove(key int) error {
	if t.IsEmpty() {
		return nil
	}
	var parent *Node[T]
	var isLeftNode bool
	current := t.root
	for current != nil {
		if current.key() == key {
			break
		}
		parent = current
		if current.data.Key() > key {
			current = current.left
			isLeftNode = true
		} else {
			current = current.right
		}
	}
	if current == nil {
		return ErrNotFoundElementByKey
	}

	if current.isLeaf() {
		if current == t.root {
			t.root = nil
		}
		if isLeftNode {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if current.left == nil {
		if current == t.root {
			t.root = current.right
		} else if isLeftNode {
			parent.left = current.right
		} else {
			parent.right = current.right
		}
		current = nil
	} else if current.right == nil {
		if current == t.root {
			t.root = current.left
		} else if isLeftNode {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
		current = nil
	} else {
		successor := fetchSuccessor(current)
		if t.root == current {
			t.root = successor
		} else if isLeftNode {
			parent.left = successor
		} else {
			parent.right = successor
		}
	}

	return nil
}

func (t *BinaryTree[T]) SymmetricTraversal(f func(T)) {
	fmt.Println("symmetric traversal")
	t.symmetricTraversal(t.root, f)
	fmt.Println()
}

func (t *BinaryTree[T]) symmetricTraversal(localRoot *Node[T], f func(T)) {
	if localRoot == nil {

	}
}

func fetchSuccessor[T Rib](node *Node[T]) *Node[T] {
	successorParent := node
	successor := node
	current := node.right
	for current != nil {
		successorParent = successor
		successor = current
		current = current.left
	}
	if successor != node.right {
		successorParent.left = successor.right
		successor.right = node.right
		successor.left = node.left
	} else {
		successor = successorParent.left
		successor.right = successorParent.right
	}
	return successor
}
