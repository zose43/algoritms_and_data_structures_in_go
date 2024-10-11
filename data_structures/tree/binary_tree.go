package tree

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

func (t *BinaryTree[T]) Remove(key int) error {
	return nil
}
